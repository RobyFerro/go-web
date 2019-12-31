package controller

import (
	"ikdev/smartcherry/exception"
	"ikdev/smartcherry/helper"
)

type AuthController struct {
	Controller
}

func (c *AuthController) Login() {

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
}
