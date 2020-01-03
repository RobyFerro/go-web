package controller

import (
	"github.com/go-redis/redis/v7"
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
	Redis    *redis.Client
}
