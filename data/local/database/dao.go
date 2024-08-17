package database

import (
	_ "github.com/mattn/go-sqlite3"
	"main.go/domain/models"
)

func CompleteLesson(progress models.Progress) {
	var existingProgress models.Progress
	DB.Where("lessons_id = ?", progress.Lesson.ID).First(&existingProgress)

	if existingProgress.Id != 0 {
		existingProgress.CurrentSpeed = progress.CurrentSpeed

		if progress.CurrentSpeed > existingProgress.BestSpeed {
			existingProgress.BestSpeed = progress.CurrentSpeed
		}

		existingProgress.Complete = progress.Complete
		DB.Save(&existingProgress)
	} else {
		DB.Create(progress)
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

func InsertPractices(practices []models.Practice) {
	for _, practice := range practices {
		DB.Save(&practice)
	}
}

func ReadPractices() []models.Practice {
	var practices []models.Practice
	DB.Find(&practices)
	return practices
}

func ReadPracticeLessons(practiceId uint) ([]models.Lesson, error) {
	var lessons []models.Lesson
	if err := DB.Preload("Content").Where("practice_id = ?", practiceId).Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func ReadLessonContent(lessonId uint) []models.LessonContent {
	var content []models.LessonContent
	DB.Where("lesson_id", lessonId).Find(&content)
	return content
}
