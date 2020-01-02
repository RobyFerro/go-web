package database

import "ikdev/go-web/database/model"

func GetModels() []interface{} {
	var models []interface{}

	models = append(models, model.User{})

	return models

}
