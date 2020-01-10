package database

import "ikdev/go-web/database/model"

// This method will return an array of models.
// Used to handle migration, seeding and drop operations.
// Every time you add a new model you should register it in this method
func GetModels() []interface{} {
	var models []interface{}

	models = append(models, model.User{})
	models = append(models, model.FailedJob{})

	return models

}
