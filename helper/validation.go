package helper

import (
	"gopkg.in/asaskevich/govalidator.v4"
)

// Validate incoming request
func ValidateRequest(data interface{}) (bool, error) {
	return govalidator.ValidateStruct(data)
}
