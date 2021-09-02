package controller

import (
	"encoding/json"
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/RobyFerro/go-web-framework/tool"
	jwt "github.com/RobyFerro/go-web/app/auth"
	"github.com/RobyFerro/go-web/app/http/validation"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserController struct {
	kernel.BaseController
}

// Insert this method will be used to insert a new user in main DB (SQL)
func (c *UserController) Insert(db *gorm.DB) {
	var data validation.NewUser
	if err := tool.DecodeJsonRequest(c.Request, &data); err != nil {
		log.Fatal(err)
	}

	if data.Password != data.RepeatPassword {
		c.Response.WriteHeader(422)
		_, _ = c.Response.Write([]byte(`{"Name":"Password","Err":{},"CustomErrorMessageExists":"false"`))
		return
	}

	// End validation
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 6)
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{
		Name:     data.Name,
		Surname:  data.Surname,
		Username: data.Username,
		Password: string(encryptedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}
}

// Profile method return information about the authenticated user.
func (c *UserController) Profile(conf *kernel.ServerConf) {
	var auth jwt.JWTAuth

	if err := auth.GetUser(c.Request, conf.Key); err != nil {
		log.Fatal(err)
	}

	jsonResponse, err := json.Marshal(auth)
	if err != nil {
		log.Fatal(err)
	}

	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.Write(jsonResponse)
}
