package typing

import (
	"fmt"
	"log"
	"main.go/data/local/database"
	"main.go/domain/models"
	"main.go/pkg/utils"
	"os"
	"time"
)

func DisplayTypingSpeed(startTime time.Time, inputWords string, lesson *models.Lesson, accuracy float64) string {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := utils.CalculateTypingSpeed(inputWords, duration)

	f, err := os.OpenFile("debugFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(fmt.Sprintf("completed the lesson at %s \n", endTime))

	progress := models.Progress{
		CurrentSpeed: currentTypingSpeed,
		BestSpeed:    currentTypingSpeed,
		Accuracy:     accuracy,
		Complete:     true,
		LessonID:     lesson.ID,
	}

	if err := database.CompleteLesson(&progress); err != nil {
		return fmt.Sprintf("Error saving your progress: %v", err)
	}

	return fmt.Sprintf("\n\n Congratulations! You have completed lesson %s\n Your typing speed is: %.2f WPM\n", lesson.Title, currentTypingSpeed)
}
