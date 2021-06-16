package app

import (
	"github.com/RobyFerro/go-web-framework/tool"
	"gopkg.in/yaml.v2"
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
}

// Configuration gets configuration struct by parsing the config.yml file.
func Configuration() (*Conf, error) {
	var conf Conf
	confFile := tool.GetDynamicPath("config.yml")
	c, err := os.Open(confFile)

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
