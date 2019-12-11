package controller

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Controller struct {
	DB       *gorm.DB
	Response http.ResponseWriter
	Request  *http.Request
}
