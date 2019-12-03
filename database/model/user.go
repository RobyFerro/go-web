package model

import (
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/exception"
)

type User struct {
	gorm.Model
	ID       int64  `gorm:"type:bigint"`
	Name     string `gorm:"type:varchar(255)"`
	Surname  string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

func (User) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Branch{})
}

func (User) Drop(db *gorm.DB) {
	if err := db.DropTableIfExists(&Branch{}).Error; err != nil {
		exception.ProcessError(err)
	}
}
