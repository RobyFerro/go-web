package helper

import (
	"encoding/json"
	"gopkg.in/asaskevich/govalidator.v4"
	"ikdev/go-web/exception"
	"net/http"
)

// Validate incoming request
func ValidateRequest(data interface{}, response http.ResponseWriter) bool {
	if valid, err := govalidator.ValidateStruct(data); valid == false {
		message, e := json.Marshal(err)

		if e != nil {
			exception.ProcessError(e)
		}

		response.WriteHeader(422)
		_, _ = response.Write(message)
		return false
	}

	return true
}
