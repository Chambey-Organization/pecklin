package typing

import (
	"fmt"
	"time"

	"main.go/pkg/models"
	"main.go/pkg/utils/typingSpeed"
)

const (
	delay = 1 * time.Second
)

func DisplayTypingSpeed(startTime time.Time, inputWords string, lessonTitle string) string {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := typingSpeed.CalculateTypingSpeed(inputWords, duration)

	var lesson models.Lesson
	lesson.CurrentSpeed = currentTypingSpeed
	lesson.BestSpeed = currentTypingSpeed
	lesson.Title = lessonTitle
	lesson.Complete = true
	//database.CompleteLesson(lesson)
	return fmt.Sprintf("\n\nCongratulations! You have completed lesson %s\nYour typing speed is: %.2f WPM\n", lessonTitle, currentTypingSpeed)
}
