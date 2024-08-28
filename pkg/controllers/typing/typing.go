package typing

import (
	"fmt"
	"main.go/data/local/database"
	"main.go/domain/models"
	"main.go/pkg/utils"
	"time"
)

func DisplayTypingSpeed(startTime time.Time, inputWords string, lesson *models.Lesson, accuracy float64) string {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := utils.CalculateTypingSpeed(inputWords, duration)

	progress := models.Progress{
		CurrentSpeed: currentTypingSpeed,
		BestSpeed:    currentTypingSpeed,
		Accuracy:     accuracy,
		Complete:     true,
		LessonID:     lesson.ID,
	}
	database.WriteToDebugFile("input string", inputWords)
	database.WriteToDebugFile("duration ", duration.String())
	database.WriteToDebugFile("input length ", fmt.Sprintf("%d", len(inputWords)))

	if err := database.CompleteLesson(&progress); err != nil {
		return fmt.Sprintf("Error saving your progress: %v", err)
	}

	return fmt.Sprintf("\n\n Congratulations! You have completed lesson %s\n Your typing speed is: %.2f WPM with an accuracy of %.2f%%\n", lesson.Title, currentTypingSpeed, accuracy)
}
