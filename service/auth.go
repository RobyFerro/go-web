package service

import (
	"ikdev/go-web/app/config"
	"ikdev/go-web/helper"
)

// Prepare Auth structure for Service Container
func SetAuth(conf config.Conf) *helper.Auth {
	return &helper.Auth{Conf: conf}
}
