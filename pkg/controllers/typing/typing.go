package typing

import (
	"fmt"
	"main.go/data/local/database"
	"main.go/domain/models"
	"time"

	"main.go/pkg/utils/typingSpeed"
)

const (
	delay = 1 * time.Second
)

func DisplayTypingSpeed(startTime time.Time, inputWords string, lessonTitle string) string {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := typingSpeed.CalculateTypingSpeed(inputWords, duration)

	database.CompleteLesson(models.Lesson{
		CurrentSpeed: currentTypingSpeed,
		BestSpeed:    currentTypingSpeed,
		Title:        lessonTitle,
		Complete:     true,
	})
	return fmt.Sprintf("\n\nCongratulations! You have completed lesson %s\nYour typing speed is: %.2f WPM\n", lessonTitle, currentTypingSpeed)
}
