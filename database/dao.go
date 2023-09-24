package database

import (
	_ "github.com/mattn/go-sqlite3"
	"main.go/pkg/models"
)

func CompleteLesson(lesson models.Lesson) {
	var existingLesson models.Lesson
	DB.Where("title = ?", lesson.Title).First(&existingLesson)

	if existingLesson.ID != 0 {
		existingLesson.CurrentSpeed = lesson.CurrentSpeed

		if lesson.CurrentSpeed > existingLesson.BestSpeed {
			existingLesson.BestSpeed = lesson.CurrentSpeed
		}

		existingLesson.Complete = lesson.Complete
		DB.Save(&existingLesson)
	} else {
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
func ReadAllLessons() []models.Lesson {
	var lessons []models.Lesson
	DB.Find(&lessons)

	return lessons
}
