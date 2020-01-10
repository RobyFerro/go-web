package model

import (
	"github.com/jinzhu/gorm"
	"ikdev/go-web/exception"
)

type FailedJob struct {
	gorm.Model
	Payload   string `gorm:"type:varchar(255)"`
	Queue     string `gorm:"type:varchar(255)"`
	Exception string `gorm:"type:varchar(255)"`
}

// Execute model migration
func (FailedJob) Migrate(db *gorm.DB) {
	db.AutoMigrate(&FailedJob{})
}

// Execute model drop
func (FailedJob) Drop(db *gorm.DB) {
	if err := db.DropTableIfExists(&FailedJob{}).Error; err != nil {
		exception.ProcessError(err)
	}
}
