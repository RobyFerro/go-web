package controller

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/helper"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController struct {
	gwf.BaseController
}

type Credentials struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// User login method.
// This method will set JWT in HTTP response
func (c *AuthController) JWTAuthentication() {
	//var user model.User
	var payload Credentials
	var user *model.User

	if err := helper.DecodeJsonRequest(c.Request, &payload); err != nil {
		gwf.ProcessError(err)
		c.Response.WriteHeader(http.StatusInternalServerError)
	}

	if valid := helper.ValidateRequest(payload, c.Response); !valid {
		c.Response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	c.Response.Header().Add("Content-Type", "application/json")
	if u, auth := attemptLogin(c.DB, &payload); !auth {
		c.Response.WriteHeader(http.StatusForbidden)
		_, _ = c.Response.Write([]byte("Invalid credentials"))
		return
	} else {
		user = u
	}
	// End check password

	// Generate JWT token
	c.BaseController.Auth.User.Name = user.Name
	c.BaseController.Auth.User.Surname = user.Surname
	c.BaseController.Auth.User.Username = user.Username
	c.BaseController.Auth.User.ID = user.ID

	if status := c.BaseController.Auth.NewToken(); !status {
		c.Response.WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response.Write([]byte(`{"error":"token_generation_failed"}`))
		return
	}

	_, _ = c.Response.Write([]byte(`{"token":"` + c.Auth.Token + `"}`))
	// End JWT token generation
	return
}

// Basic authentication method
// Todo: complete this method to provide basic authentication
func (c *AuthController) BasicAuthentication() {
	var payload Credentials
	if err := helper.DecodeJsonRequest(c.Request, &payload); err != nil {
		gwf.ProcessError(err)
		c.Response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if valid := helper.ValidateRequest(payload, c.Response); !valid {
		c.Response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if user, auth := attemptLogin(c.DB, &payload); !auth {
		c.Response.WriteHeader(http.StatusForbidden)
		_, _ = c.Response.Write([]byte("Invalid credentials"))
		return
	} else {
		if err := createAuthSession(c.Session, user, c.Request, c.Response); err != nil {
			c.Response.WriteHeader(http.StatusInternalServerError)
			_, _ = c.Response.Write([]byte("Invalid credentials"))
			return
		}
	}

	c.Response.WriteHeader(http.StatusAccepted)
}

// Create auth session for the current user
func createAuthSession(s *sessions.CookieStore, user *model.User, r *http.Request, w http.ResponseWriter) error {
	if session, err := s.Get(r, "basic-auth"); err != nil {
		return err
	} else {
		// Remove password from user structure
		//userJson, _ := json.Marshal(user)
		session.Values["user"] = user
		if err := session.Save(r, w); err != nil {
			return err
		}
	}

	return nil
}

// Attempt login
func attemptLogin(db *gorm.DB, cred *Credentials) (*model.User, bool) {
	var user model.User
	if err := db.Where("username = ?", cred.Username).Find(&user); err != nil && err.RecordNotFound() {
		return nil, false
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		return nil, false
	}

	return &user, true
}
