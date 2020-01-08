package job

import (
	"encoding/json"
	"fmt"
	"ikdev/go-web/config"
	"ikdev/go-web/database"
)

type MailStruct struct {
	To      []string `json:"to"`
	Message string   `json:"message"`
}

func (Job) Mail(payload string) {
	var data MailStruct

	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		redis := database.ConnectRedis(config.Configuration())
		Job{}.PutOnFailed(payload, redis)
	}

	fmt.Println("Mail sent!")
}
