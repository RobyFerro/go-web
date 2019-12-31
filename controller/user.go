package controller

import (
	"golang.org/x/crypto/bcrypt"
	"ikdev/smartcherry/database/model"
	"ikdev/smartcherry/exception"
	"ikdev/smartcherry/helper"
)

type UserController struct {
	Controller
}

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
