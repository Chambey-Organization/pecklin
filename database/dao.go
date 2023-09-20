package database

import (
	_ "github.com/mattn/go-sqlite3"
	"main.go/pkg/models"
)

func CompleteLesson(lesson models.Lesson) {
	DB.FirstOrCreate(&lesson, models.Lesson{Title: lesson.Title})
	DB.Model(&lesson).Update(models.Lesson{CurrentSpeed: lesson.CurrentSpeed})

	if lesson.CurrentSpeed > lesson.BestSpeed {
		lesson.BestSpeed = lesson.CurrentSpeed
		DB.Model(&lesson).Update("BestSpeed", lesson.BestSpeed)
	}
}

func RedoLessons() {
	DB.Model(&models.Lesson{}).Where("complete = ?", true).Update("complete", false)
}

func ReadCompletedLesson() []models.Lesson {
	var lessons []models.Lesson
	DB.Where("complete = ?", false).Find(&lessons)
	return lessons
}

func ReadALessonData(searchLesson models.Lesson) models.Lesson {
	var lesson models.Lesson

	// Retrieve a lesson by title
	DB.Where("lesson = ?", searchLesson.Title).First(&lesson)

	return lesson
}
