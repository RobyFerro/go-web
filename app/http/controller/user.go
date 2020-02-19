package controller

import (
	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/helper"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	gwf.BaseController
}

// This method will be used to insert a new user in main DB (SQL)
func (c *UserController) Insert(db *gorm.DB) {

	type NewUser struct {
		Name           string `json:"name" valid:"required,alpha"`
		Surname        string `json:"surname" valid:"required,alpha"`
		Username       string `json:"username" valid:"required,alpha"`
		Password       string `json:"password" valid:"required,alpha"`
		RepeatPassword string `json:"repeat-password" valid:"required,alpha"`
	}

	var data NewUser
	if err := helper.DecodeJsonRequest(c.Request, &data); err != nil {
		gwf.ProcessError(err)
	}

	// Validation
	if valid := helper.ValidateRequest(data, c.Response); valid == false {
		return
	}

	if data.Password != data.RepeatPassword {
		c.Response.WriteHeader(422)
		_, _ = c.Response.Write([]byte(`{"Name":"Password","Err":{},"CustomErrorMessageExists":"false"`))
		return
	}
	// End validation

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		gwf.ProcessError(err)
	}

	user := model.User{
		Name:     data.Name,
		Surname:  data.Surname,
		Username: data.Username,
		Password: string(encryptedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		gwf.ProcessError(err)
	}

	c.Response.WriteHeader(200)
}

// Profile method return information about the authenticated user.
func (c *UserController) Profile(conf *gwf.Conf) {
	var auth gwf.Auth
	if err := auth.GetUser(c.Request, conf.App.Key); err != nil {
		gwf.ProcessError(err)
	}

	_, _ = c.Response.Write([]byte("Authorization ok"))
}
