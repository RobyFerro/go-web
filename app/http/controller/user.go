package controller

import (
	"github.com/RobyFerro/go-web/database/model"
	"github.com/RobyFerro/go-web/exception"
	"github.com/RobyFerro/go-web/helper"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	BaseController
}

// This method will be used to insert a new user in main DB (SQL)
func (c *UserController) Insert() {

	type NewUser struct {
		Name           string `json:"name" valid:"required,alpha"`
		Surname        string `json:"surname" valid:"required,alpha"`
		Username       string `json:"username" valid:"required,alpha"`
		Password       string `json:"password" valid:"required,alpha"`
		RepeatPassword string `json:"repeat-password" valid:"required,alpha"`
	}

	var data NewUser
	if err := helper.DecodeJsonRequest(c.Request, &data); err != nil {
		exception.ProcessError(err)
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
		exception.ProcessError(err)
	}

	user := model.User{
		Name:     data.Name,
		Surname:  data.Surname,
		Username: data.Username,
		Password: string(encryptedPassword),
	}

	if err := c.DB.Create(&user).Error; err != nil {
		exception.ProcessError(err)
	}

	c.Response.WriteHeader(200)
}

//Used to check authenticated user
func (c *UserController) Profile() {
	c.Auth.GetUser(c.Request)

	_, _ = c.Response.Write([]byte("Authorization ok"))
}
