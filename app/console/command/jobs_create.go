package command

import (
	"fmt"
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/app/kernel"
	"github.com/RobyFerro/go-web/exception"
	"io/ioutil"
	"strings"
)

type JobCreate struct {
	Signature   string
	Description string
}

// Command registration
func (c *JobCreate) Register() {
	c.Signature = "job:create <name>"      // Change command signature
	c.Description = "Create new async job" // Change command description
}

// Command business logic
func (c *JobCreate) Run(kernel *kernel.HttpKernel, args string, console map[string]interface{}) {
	splitName := strings.Split(strings.ToLower(args), "_")
	for i, name := range splitName {
		splitName[i] = strings.Title(name)
	}

	cName := strings.Join(splitName, "")
	input, _ := ioutil.ReadFile(config.GetDynamicPath("raw/job.raw"))

	cContent := strings.ReplaceAll(string(input), "@@TMP@@", cName)
	cFile := fmt.Sprintf("%s/%s.go", config.GetDynamicPath("job"), strings.ToLower(args))
	if err := ioutil.WriteFile(cFile, []byte(cContent), 0755); err != nil {
		exception.ProcessError(err)
	}

	fmt.Printf("\nSUCCESS: Your %s job has been created at %s\n", cName, cFile)
}
