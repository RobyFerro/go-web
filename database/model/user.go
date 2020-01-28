package model

import (
	"github.com/brianvoe/gofakeit/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"ikdev/go-web/exception"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"`
	Surname  string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}

// Execute model migration
// Deprecated: this method has been replaced by global .sql migration system
func (User) Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

// Execute model drop
// Deprecated: this method has been replaced by global .sql migration system
func (User) Drop(db *gorm.DB) {
	if err := db.DropTableIfExists(&User{}).Error; err != nil {
		exception.ProcessError(err)
	}
}

// Execute model seeding
func (User) Seed(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		password := gofakeit.Password(true, true, true, true, false, 32)
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

		user := User{
			Name:     gofakeit.FirstName(),
			Surname:  gofakeit.LastName(),
			Username: gofakeit.Username(),
			Password: string(encryptedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			exception.ProcessError(err)
		}
	}
}
