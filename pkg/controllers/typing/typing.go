package typing

import (
	"fmt"
	"main.go/data/local/database"
	"main.go/domain/models"
	"time"

	"main.go/pkg/utils/typingSpeed"
)

func DisplayTypingSpeed(startTime time.Time, inputWords string, lesson *models.Lesson, accuracy float64) string {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := typingSpeed.CalculateTypingSpeed(inputWords, duration)

	progress := models.Progress{
		CurrentSpeed: currentTypingSpeed,
		BestSpeed:    currentTypingSpeed,
		Accuracy:     accuracy,
		Complete:     true,
		LessonID:     lesson.ID,
	}

	if err := database.CompleteLesson(progress); err != nil {
		return fmt.Sprintf("Error saving your progress: %v", err)
	}

	return fmt.Sprintf("\n\n Congratulations! You have completed lesson %s\nYour typing speed is: %.2f WPM\n", lesson.Title, currentTypingSpeed)
}
