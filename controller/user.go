package controller

import (
	"golang.org/x/crypto/bcrypt"
	"ikdev/smartcherry/database/model"
	"ikdev/smartcherry/exception"
)

type UserController struct {
	Controller
}

func (c *UserController) Insert() {
	_ = c.Request.ParseForm()
	password, _ := bcrypt.GenerateFromPassword([]byte(c.Request.Form.Get("password")), 14)

	user := model.User{
		Name:     c.Request.Form.Get("name"),
		Surname:  c.Request.Form.Get("surname"),
		Username: c.Request.Form.Get("username"),
		Password: string(password),
	}

	if err := c.DB.Create(&user).Error; err != nil {
		exception.ProcessError(err)
	}
}
