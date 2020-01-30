package command

import (
	"github.com/olekukonko/tablewriter"
	"ikdev/go-web/app"
	"os"
	"reflect"
)

type ShowCommands struct {
	Signature   string
	Description string
}

func (c *ShowCommands) Register() {
	c.Signature = "show:commands"
	c.Description = "Show Go-Web commands list"
}

func (c *ShowCommands) Run(sc *app.HttpKernel, args string) {
	var data [][]string

	for _, c := range Register {
		m := reflect.ValueOf(c).MethodByName("Register")
		m.Call([]reflect.Value{})

		cmd := reflect.ValueOf(c).Elem()

		signature := cmd.FieldByName("Signature").String()
		description := cmd.FieldByName("Description").String()

		data = append(data, []string{signature, description})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"COMMAND", "DESCRIPTION"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
