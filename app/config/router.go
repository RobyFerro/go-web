package config

import (
	"github.com/RobyFerro/go-web/exception"
	"gopkg.in/yaml.v2"
	"os"
)

// Structure used to decode all route presents into routing.yml file.
type Route struct {
	Path        string   `yaml:"path"`
	Action      string   `yaml:"action"`
	Method      string   `yaml:"method"`
	Description string   `yaml:"description"`
	Middleware  []string `yaml:"middleware"`
	Prefix      string   `yaml:"prefix"`
}

// Structure used to decode all groups presents into the routing.yml file.
type Group struct {
	Prefix     string `yaml:"prefix"`
	Routes     map[string]Route
	Middleware []string
}

// Main structure of web router.
type Router struct {
	Routes map[string]Route `yaml:"routes"`
	Groups map[string]Group `yaml:"groups"`
}

// Parse router.yml file (present in Go-Web root dir) and return a Router structure.
// This structure will be used by the HTTP kernel to setup every routes.
func ConfigurationWeb() Router {
	var conf Router
	routePath := GetDynamicPath("routing.yml")
	c, err := os.Open(routePath)

	if err != nil {
		exception.ProcessError(err)
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(&conf); err != nil {
		exception.ProcessError(err)
	}

	return conf
}
