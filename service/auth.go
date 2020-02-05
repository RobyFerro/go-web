package service

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/helper"
)

// Prepare Auth structure for Service Container
func SetAuth(conf config.Conf) *helper.Auth {
	return &helper.Auth{Conf: conf}
}
