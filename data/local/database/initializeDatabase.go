package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
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
