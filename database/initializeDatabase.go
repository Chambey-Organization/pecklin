package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"main.go/pkg/models"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error
	DB, err = gorm.Open("sqlite3", "pecklin.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	// Create the "lessons" table if it does not exist
	DB.AutoMigrate(&models.Lesson{})
}
