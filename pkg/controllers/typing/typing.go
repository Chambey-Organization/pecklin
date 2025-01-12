package typing

import (
	"main.go/data/local/database"
	"main.go/domain/models"
	"main.go/pkg/utils"
	"time"
)

func SaveTypingSpeed(startTime time.Time, inputWords string, lesson *models.Lesson, accuracy float64) error {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := utils.CalculateTypingSpeed(inputWords, duration)

	progressLesson := models.Lesson{
		ID:         lesson.ID,
		PracticeID: lesson.PracticeID,
		Title:      lesson.Title,
		Active:     lesson.Active,
		Content:    lesson.Content,
	}

	progress := models.Progress{
		CurrentSpeed: currentTypingSpeed,
		BestSpeed:    currentTypingSpeed,
		Accuracy:     accuracy,
		Complete:     true,
		Lesson:       progressLesson,
		LessonID:     lesson.ID,
	}

	database.WriteToDebugFile("saved the lesson progress", "")

	return database.CompleteLesson(&progress)
}
