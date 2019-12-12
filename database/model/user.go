package model

import (
	"github.com/brianvoe/gofakeit/v4"
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/exception"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Surname  string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

func (User) Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func (User) Drop(db *gorm.DB) {
	if err := db.DropTableIfExists(&Branch{}).Error; err != nil {
		exception.ProcessError(err)
	}
}

func (User) Seed(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		user := User{
			Name:     gofakeit.FirstName(),
			Surname:  gofakeit.LastName(),
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 32),
		}

		if err := db.Create(&user).Error; err != nil {
			exception.ProcessError(err)
		}
	}
}
