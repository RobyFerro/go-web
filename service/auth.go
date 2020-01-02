package service

import (
	"ikdev/go-web/config"
	"ikdev/go-web/helper"
)

func SetAuth(conf config.Conf) *helper.Auth {
	return &helper.Auth{Conf: conf}
}
