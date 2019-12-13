package helper

import (
	"gopkg.in/asaskevich/govalidator.v4"
)

// Validate incoming request
func ValidateRequest(data interface{}) (bool, error) {
	result, err := govalidator.ValidateStruct(data)
	if err != nil {
		return false, err
	}

	return result, nil
}
