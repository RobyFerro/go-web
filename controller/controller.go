package controller

import (
	"github.com/jinzhu/gorm"
	"ikdev/go-web/config"
	"ikdev/go-web/helper"
	"net/http"
)

type Controller struct {
	DB       *gorm.DB
	Response http.ResponseWriter
	Request  *http.Request
	Config   config.Conf
	Auth     *helper.Auth
}
