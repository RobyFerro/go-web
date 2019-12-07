package database

import "ikdev/smartcherry/database/model"

func GetModels() []interface{} {
	var models []interface{}

	models = append(models, model.Branch{})
	models = append(models, model.User{})

	return models

}
