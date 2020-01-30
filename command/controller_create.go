package command

import (
	"fmt"
	"go.uber.org/dig"
	"ikdev/go-web/config"
	"ikdev/go-web/exception"
	"io/ioutil"
	"strings"
)

type ControllerCreate struct {
	Signature   string
	Description string
}

func (c *ControllerCreate) Register() {
	c.Signature = "controller:create <name>"
	c.Description = "Create new controller"
}

func (c *ControllerCreate) Run(sc *dig.Container, args string) {
	cName := strings.Title(strings.ToLower(args))
	input, _ := ioutil.ReadFile(config.GetFilePath("raw/controller.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", config.GetFilePath("controller"), strings.ToLower(args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		exception.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %sController has been created at %s", cName, cFile)
	fmt.Printf("Do not forget to register it!")
}
