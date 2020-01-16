package config

import (
	"gopkg.in/yaml.v2"
	"ikdev/go-web/exception"
	"os"
)

type Conf struct {
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Mongo struct {
		Database string `yaml:"database"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"mongodb"`
	Elastic struct {
		Hosts []string `yaml:"hosts"`
	} `yaml:"elasticsearch"`
	Server struct {
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
		Ssl      bool   `yaml:"ssl"`
		SslCert  string `yaml:"sslcert"`
		SslKey   string `yaml:"sslkey"`
		RunUser  string `yaml:"run-user"`
		RunGroup string `yaml:"run-group"`
	} `yaml:"server"`
	App struct {
		Key string `yaml:"key"`
	} `yaml:"app"`
}

// Get configuration struct
func Configuration() Conf {
	var conf Conf
	getConf(&conf)

	return conf
}

// Parse configuration .yml file in struct
func getConf(conf *Conf) {
	confFile := GetFilePath("config.yml")
	c, err := os.Open(confFile)

	if err != nil {
		exception.ProcessError(err)
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(conf); err != nil {
		exception.ProcessError(err)
	}
}
