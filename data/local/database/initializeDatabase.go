package database

import (
	"github.com/jinzhu/gorm"
	"main.go/domain/models"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error
	DB, err = gorm.Open("sqlite3", "pecklin.db")
	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&models.Lesson{})
	DB.AutoMigrate(&models.Progress{})
	DB.AutoMigrate(&models.Practice{})
	DB.AutoMigrate(&models.LessonContent{})
}
