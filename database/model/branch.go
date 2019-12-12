package model

import (
	"github.com/brianvoe/gofakeit/v4"
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/exception"
)

type Branch struct {
	gorm.Model
	UserID  uint `gorm:"type:int"`
	User    User
	Name    string `gorm:"type:varchar(255)"`
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

func (Branch) Seed(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		var user User

		if err := db.Take(&user).Error; err != nil {
			exception.ProcessError(err)
		}

		branch := Branch{
			UserID:  user.ID,
			Name:    gofakeit.Name(),
			Content: gofakeit.Paragraph(10, 20, 20, "."),
		}

		if err := db.Create(&branch).Error; err != nil {
			exception.ProcessError(err)
		}
	}
}
