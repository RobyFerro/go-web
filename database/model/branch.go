package model

import "github.com/jinzhu/gorm"

type Branch struct {
	gorm.Model
	UserID  int64 `gorm:"type:bigint"`
	User    User
	ID      int64  `gorm:"type:bigint"`
	Name    string `gorm:"type:bigint"`
	Content string `gorm:"type:text"`
}
