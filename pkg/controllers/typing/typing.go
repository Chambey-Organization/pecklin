package typing

import (
	"fmt"
	"main.go/domain/models"
	"time"

	"main.go/pkg/utils/typingSpeed"
)

const (
	delay = 1 * time.Second
)

func DisplayTypingSpeed(startTime time.Time, inputWords string, lesson models.Lesson, accuracy float64) string {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := typingSpeed.CalculateTypingSpeed(inputWords, duration)

	/* database.CompleteLesson(
	models.Progress{
		CurrentSpeed: currentTypingSpeed,
		BestSpeed:    currentTypingSpeed,
		Accuracy:     accuracy,
		Complete:     true,
		Lesson:       lesson,
	}) */
	return fmt.Sprintf("\n\nCongratulations! You have completed lesson %s\nYour typing speed is: %.2f WPM\n", lesson.Title, currentTypingSpeed)
}
