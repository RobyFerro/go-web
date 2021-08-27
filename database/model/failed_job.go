package model

import (
	"github.com/jinzhu/gorm"
)

type FailedJob struct {
	gorm.Model
	Payload   string `gorm:"type:varchar(255)"`
	Queue     string `gorm:"type:varchar(255)"`
	Exception string `gorm:"type:varchar(255)"`
}
