package model

import (
	"github.com/brianvoe/gofakeit/v4"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Surname  string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

// Seed executes seeding in defined table
func (User) Seed(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

		user := User{
			Name:     gofakeit.FirstName(),
			Surname:  gofakeit.LastName(),
			Username: gofakeit.Username(),
			Password: string(encryptedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			log.Error(err)
		}
	}
}
