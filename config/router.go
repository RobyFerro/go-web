package config

import (
	"gopkg.in/yaml.v2"
	"ikdev/go-web/exception"
	"os"
)

type Route struct {
	Path       string `yaml:"path"`
	Action     string `yaml:"action"`
	Method     string `yaml:"method"`
	Middleware string `yaml:"middleware"`
	Prefix     string `yaml:"prefix"`
}

type Group struct {
	Prefix     string `yaml:"prefix"`
	Routes     map[string]Route
	Middleware []string
}

type Router struct {
	Routes map[string]Route `yaml:"routes"`
	Groups map[string]Group `yaml:"groups"`
}

// Get routing struct
func ConfigurationWeb() Router {
	var conf Router
	getConfWeb(&conf)

	return conf
}

// Parse routing.yml file in struct
func getConfWeb(conf *Router) {
	c, err := os.Open("routing.yml")

	if err != nil {
		exception.ProcessError(err)
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(conf); err != nil {
		exception.ProcessError(err)
	}
}
