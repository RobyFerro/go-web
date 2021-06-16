package model

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type FailedJob struct {
	gorm.Model
	Payload   string `gorm:"type:varchar(255)"`
	Queue     string `gorm:"type:varchar(255)"`
	Exception string `gorm:"type:varchar(255)"`
}

// Execute model migration
// Deprecated: this method has been replaced by global .sql migration system
func (FailedJob) Migrate(db *gorm.DB) {
	db.AutoMigrate(&FailedJob{})
}

// Execute model drop
// Deprecated: this method has been replaced by global .sql migration system
func (FailedJob) Drop(db *gorm.DB) {
	if err := db.DropTableIfExists(&FailedJob{}).Error; err != nil {
		log.Error(err)
	}
}
