package job

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type MailStruct struct {
	To      []string `json:"to"`
	Message string   `json:"message"`
}

// Example job
// Every job must return a tuple (bool, error)
func (Job) Mail(payload string) (bool, error) {
	var data MailStruct

	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return false, errors.New("Cannot unmarshal payload")
	}

	return true, nil
}
