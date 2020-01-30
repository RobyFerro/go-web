package job

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"ikdev/go-web/app/config"
	"ikdev/go-web/exception"
	"net/smtp"
)

type MailStruct struct {
	To      []string `json:"to"`
	Message string   `json:"message"`
}

// Execute send mail
func (Job) Mail(payload string) (bool, error) {
	var data MailStruct
	conf := config.Configuration()

	auth := smtp.PlainAuth("", conf.Mail.Host, conf.Mail.Password, conf.Mail.Host)
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return false, errors.New("Cannot unmarshal payload")
	}

	server := fmt.Sprintf("%s:%d", conf.Mail.Host, conf.Mail.Port)
	if err := smtp.SendMail(server, auth, conf.Mail.From, data.To, []byte(data.Message)); err != nil {
		exception.ProcessError(err)
	}

	return true, nil
}
