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
}

func (c *Auth) GetUser(req *http.Request, conf config.Conf) {
	bearerSchema := "Bearer "
	tokenString := req.Header.Get("Authorization")

	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(tokenString[len(bearerSchema):], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.App.Key), nil
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

func (c *Auth) NewToken(key string) bool {
	c.User.Password = ""
	userDataString, _ := json.Marshal(c.User)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userDataString),
		"exp":  time.Now().Add(time.Hour * time.Duration(2)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(key))
	c.Token = tokenString

	if err != nil {
		return false
	}

	return true
}
