package command

import (
	"fmt"
	"ikdev/go-web/app/config"
	"ikdev/go-web/app/kernel"
	"ikdev/go-web/exception"
	"io/ioutil"
	"strings"
)

type CmdCreate struct {
	Signature   string
	Description string
}

func (c *CmdCreate) Register() {
	c.Signature = "cmd:create <name>"
	c.Description = "Create new command"
}

func (c *CmdCreate) Run(kernel *kernel.HttpKernel, args string, console map[string]interface{}) {

	splitName := strings.Split(strings.ToLower(args), "_")
	for i, name := range splitName {
		splitName[i] = strings.Title(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := ioutil.ReadFile(config.GetFilePath("raw/command.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", config.GetFilePath("command"), strings.ToLower(args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		exception.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %s command has been created at %s\n", cName, cFile)
	fmt.Printf("Do not forget to register it!")
}
