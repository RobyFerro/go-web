package exception

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type ErrorConfiguration struct {
	Exception struct {
		Sentry string `yaml:"sentry"`
	} `yaml:"exception"`
}

var configuration ErrorConfiguration

// Generic method to handle errors.
// You can customize this method to implement your error handling.
// Es.: You can implement "Sentry" or other error tracking system
func ProcessError(err error) {
	configuration = getExceptionConfig()

	if configuration.Exception.Sentry != "" {
		SentryReport(err)
	}

	fmt.Println(err.Error())
	Log(err.Error())
}

// Obtain exception configuration from global app settings
func getExceptionConfig() ErrorConfiguration {
	var data ErrorConfiguration
	c, err := os.Open("config.yml")

	if err != nil {
		fmt.Println(err.Error())
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(&data); err != nil {
		fmt.Println(err.Error())
	}

	return data
}
