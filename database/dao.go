package database

import (
	_ "github.com/mattn/go-sqlite3"
	"main.go/pkg/models"
)

func CompleteLesson(lesson models.Lesson) {
	var existingLesson models.Lesson
	DB.Where("title = ?", lesson.Title).First(&existingLesson)

	// Check if the record exists
	if existingLesson.ID != 0 {
		// Update the existing record with new data
		existingLesson.CurrentSpeed = lesson.CurrentSpeed

		if lesson.CurrentSpeed > existingLesson.BestSpeed {
			existingLesson.BestSpeed = lesson.CurrentSpeed
		}

		existingLesson.Complete = lesson.Complete

		// Save the updated record back to the database
		DB.Save(&existingLesson)
	} else {
		// If the record doesn't exist, create a new one
		DB.Create(&lesson)
	}
}

func RedoLessons() {
	DB.Model(&models.Lesson{}).Where("complete = ?", true).Update("complete", false)
}

func ReadCompletedLesson() []models.Lesson {
	var lessons []models.Lesson
	DB.Where("complete = ?", true).Find(&lessons)

	return lessons
}

func ReadALessonData(searchLesson models.Lesson) models.Lesson {
	var lesson models.Lesson
	DB.Where("lesson = ?", searchLesson.Title).First(&lesson)

	return lesson
}
