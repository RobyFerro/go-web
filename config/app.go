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
	Server struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	App struct {
		Key string `yaml:"key"`
	} `yaml:"app"`
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
