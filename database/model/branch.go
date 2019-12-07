package model

import (
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/exception"
)

type Branch struct {
	gorm.Model
	UserID  int64 `gorm:"type:bigint"`
	User    User
	Name    string `gorm:"type:bigint"`
	Content string `gorm:"type:text"`
}

func (Branch) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Branch{})
}

func (Branch) Drop(db *gorm.DB) {
	if err := db.DropTableIfExists(&Branch{}).Error; err != nil {
		exception.ProcessError(err)
	}
}
