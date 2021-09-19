package controller

import (
	"github.com/RobyFerro/go-web/app/auth"
	"net/http"
	"time"

	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	kernel.BaseController
}

// JWTAuthentication provides user authentication with JWT
func (c *AuthController) JWTAuthentication(db *gorm.DB, conf *kernel.ServerConf, req kernel.Request) {
	var user *model.User
	var jwt auth.JWTAuth

	c.Response.Header().Add("Content-Type", "application/json")
	if u, ok := attemptLogin(db, req["username"].(string), req["password"].(string)); !ok {
		c.Response.WriteHeader(http.StatusForbidden)
		_, _ = c.Response.Write([]byte("Invalid credentials"))
		return
	} else {
		user = u
	}

	jwt.Name = user.Name
	jwt.Surname = user.Surname
	jwt.Username = user.Username
	jwt.ID = user.ID

	if token, ok := jwt.NewToken(conf.Key, 2*time.Hour); !ok {
		c.Response.WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response.Write([]byte(`{"error":"token_generation_failed"}`))
		return
	} else {
		_, _ = c.Response.Write([]byte(`{"token":"` + token + `"}`))
	}

	return
}

// BasicAuthentication perform basic authentication method
func (c *AuthController) BasicAuthentication(db *gorm.DB, session *sessions.CookieStore, req kernel.Request) {
	if user, ok := attemptLogin(db, req["username"].(string), req["password"].(string)); !ok {
		c.Response.WriteHeader(http.StatusForbidden)
		_, _ = c.Response.Write([]byte("Invalid credentials"))
		return
	} else {
		if err := createAuthSession(session, user, c.Request, c.Response); err != nil {
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
func attemptLogin(db *gorm.DB, username, password string) (*model.User, bool) {
	var user model.User
	if err := db.Where("username = ?", username).Find(&user); err != nil && err.RecordNotFound() {
		return nil, false
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, false
	}

	return &user, true
}
