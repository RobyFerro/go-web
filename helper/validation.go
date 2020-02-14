package helper

import (
	"encoding/json"
	"github.com/RobyFerro/go-web-framework"
	"gopkg.in/asaskevich/govalidator.v4"
	"net/http"
)

// Validate incoming request
func ValidateRequest(data interface{}, response http.ResponseWriter) bool {
	if valid, err := govalidator.ValidateStruct(data); valid == false {
		message, e := json.Marshal(err)

		if e != nil {
			gwf.ProcessError(e)
		}

		response.WriteHeader(422)
		_, _ = response.Write(message)
		return false
	}

	return true
}
