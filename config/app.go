package config

import (
	"gopkg.in/yaml.v2"
	"ikdev/smartcherry/exception"
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
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func Configuration() Conf {
	var conf Conf
	getConf(&conf)

	return conf
}

func getConf(conf *Conf) {
	c, err := os.Open("config.yml")

	if err != nil {
		exception.ProcessError(err)
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(conf); err != nil {
		exception.ProcessError(err)
	}
}
