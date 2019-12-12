package database

import "ikdev/smartcherry/database/model"

func GetModels() []interface{} {
	var models []interface{}

	models = append(models, model.User{})
	models = append(models, model.Branch{})

	return models

}
