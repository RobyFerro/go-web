package helper

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"ikdev/go-web/config"
	"ikdev/go-web/database/model"
	"ikdev/go-web/exception"
	"net/http"
	"time"
)

type Auth struct {
	User  model.User `json:"user"`
	Token string
	Conf  config.Conf
}

// Get user struct from authentication token (JWT)
func (c *Auth) GetUser(req *http.Request) {
	bearerSchema := "Bearer "
	tokenString := req.Header.Get("Authorization")

	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenString[len(bearerSchema):], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.Conf.App.Key), nil
	})

	for key, val := range claims {
		if key == "user" {
			userData := val.(string)
			err := json.Unmarshal([]byte(userData), &c.User)

			if err != nil {
				exception.ProcessError(err)
			}
		}
	}
}

// Issue new JWT token
func (c *Auth) NewToken() bool {
	c.User.Password = ""
	userDataString, _ := json.Marshal(c.User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  time.Now().Add(time.Hour * time.Duration(2)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(c.Conf.App.Key))
	c.Token = tokenString

	if err != nil {
		return false
	}

	return true
}

// Refresh JWT token
func (c *Auth) RefreshToken() bool {
	expirationTime := time.Now().Add(5 * time.Minute)
	userDataString, _ := json.Marshal(c.User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  expirationTime,
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(c.Conf.App.Key))
	c.Token = tokenString

	if err != nil {
		return false
	}

	return true
}
