package command

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework"
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

func (c *CmdCreate) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {

	splitName := strings.Split(strings.ToLower(args), "_")
	for i, name := range splitName {
		splitName[i] = strings.Title(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := ioutil.ReadFile(gwf.GetDynamicPath("raw/command.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", gwf.GetDynamicPath("app/console/command"), strings.ToLower(args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		gwf.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %s command has been created at %s\n", cName, cFile)
	fmt.Printf("Do not forget to register it!")
}
