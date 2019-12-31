package controller

import (
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/config"
	"net/http"
)

type Controller struct {
	DB       *gorm.DB
	Response http.ResponseWriter
	Request  *http.Request
	Config   config.Conf
}
