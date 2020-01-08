package job

import (
	"encoding/json"
	"fmt"
	"ikdev/go-web/exception"
)

type MailStruct struct {
	To      []string `json:"to"`
	Message string   `json:"message"`
}

func (Job) Mail(payload string) {
	var data MailStruct
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		exception.ProcessError(err)
	}

	fmt.Println("Mail sent!")
}
