package database

import "github.com/RobyFerro/go-web/database/model"

// This method will return an array of models.
// Used to handle migration, seeding and drop operations.
// Every time you add a new model you should register it in this method
func GetModels() []interface{} {
	return []interface{}{
		model.User{},
		model.FailedJob{},
		// Here is where you've to register your custom models
	}
}
