package controller

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"ikdev/go-web/database/model"
	"ikdev/go-web/exception"
	"ikdev/go-web/helper"
	"net/http"
)

type AuthController struct {
	BaseController
}

type Credentials struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// User login method.
// This method will set JWT in HTTP response
func (c *AuthController) Login() {
	var user model.User
	var payload Credentials

	if err := helper.DecodeJsonRequest(c.Request, &payload); err != nil {
		exception.ProcessError(err)
	}

	if valid := helper.ValidateRequest(payload, c.Response); !valid {
		return
	}

	c.Response.Header().Add("Content-Type", "application/json")
	if auth := attemptLogin(c.DB, &payload); !auth {
		c.Response.WriteHeader(http.StatusForbidden)
		_, _ = c.Response.Write([]byte("Invalid credentials"))
		return
	}
	// End check password

	// Generate JWT token
	c.BaseController.Auth.User = user
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
		exception.ProcessError(err)
	}

	if valid := helper.ValidateRequest(payload, c.Response); !valid {
		return
	}

	if auth := attemptLogin(c.DB, &payload); !auth {
		// login failed
	}

	// login successful
}

// Attempt login
func attemptLogin(db *gorm.DB, cred *Credentials) bool {
	var user model.User

	if err := db.Where("username = ?", cred.Username).Find(&user); err != nil {
		if err.RecordNotFound() {
			return false
		}
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return false
	}

	return true
}
