package config

import "github.com/RobyFerro/go-web-framework/kernel"

func GetSever() kernel.ServerConf {
	return kernel.ServerConf{
		Port:    8005,
		SSL:     false,
		SSLCert: "storage/certs/tls.crt",
		SSLKey:  "storage/certs/tls.key",
		Key:     "REPLACE_WITH_CUSTOM_APP_KEY",
	}
}
