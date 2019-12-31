package controller

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"ikdev/smartcherry/database/model"
	"ikdev/smartcherry/exception"
	"ikdev/smartcherry/helper"
	"net/http"
	"time"
)

type AuthController struct {
	Controller
}

const (
	appKey = "golangcode.com"
)

func (c *AuthController) Login() {
	var user model.User

	type AuthRequest struct {
		Username string `json:"username" valid:"required"`
		Password string `json:"password" valid:"required"`
	}

	var payload AuthRequest
	if err := helper.DecodeJsonRequest(c.Request, &payload); err != nil {
		exception.ProcessError(err)
	}

	if valid := helper.ValidateRequest(payload, c.Response); valid == false {
		return
	}

	c.Response.Header().Add("Content-Type", "application/json")
	if err := c.DB.Where("username = ?", payload.Username).Find(&user); err != nil {
		if err.RecordNotFound() {
			c.Response.WriteHeader(403)
			_, _ = c.Response.Write([]byte("Invalid credentials"))
			return
		}
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		c.Response.WriteHeader(403)
		_, _ = c.Response.Write([]byte("Invalid credentials"))
		return
	}
	// End check password

	// issue JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.Username,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(appKey))
	if err != nil {
		c.Response.WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response.Write([]byte(`{"error":"token_generation_failed"}`))
		return
	}
	_, _ = c.Response.Write([]byte(`{"token":"` + tokenString + `"}`))
	return
	// end issue JWT
}
