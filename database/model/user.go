package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       int64  `gorm:"type:bigint"`
	Name     string `gorm:"type:varchar(255)"`
	Surname  string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
