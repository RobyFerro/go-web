package model

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/jinzhu/gorm"
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
		gwf.ProcessError(err)
	}
}
