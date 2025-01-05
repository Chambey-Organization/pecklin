package database

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main.go/domain/models"
	"os"
)

func CompleteLesson(progress *models.Progress) error {
	var existingProgress models.Progress
	DB.Where("lesson_id = ?", progress.LessonID).First(&existingProgress)

	if existingProgress.Id != 0 {
		existingProgress.CurrentSpeed = progress.CurrentSpeed

		if progress.CurrentSpeed > existingProgress.BestSpeed {
			existingProgress.BestSpeed = progress.CurrentSpeed
		}

		existingProgress.Complete = progress.Complete
		result := DB.Save(&existingProgress)
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := DB.Create(progress)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func GetResults() []models.Progress {
	var allProgress []models.Progress
	DB.Preload("Lesson").Find(&allProgress)
	return allProgress
}

func GetLessonResultResult(id uint) (models.Progress, error) {
	var result models.Progress
	err := DB.Preload("Lesson").Where("lesson_id = ?", id).Find(&result)
	return result, err.Error
}

func InsertPractices(practices []models.Practice) {
	if len(practices) > 0 {
		DB.Delete(&models.Practice{})
		DB.Delete(&models.Lesson{})
		DB.Delete(&models.LessonContent{})
	}

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

func WriteToDebugFile(description string, input string) {
	f, err := os.OpenFile(".debugFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(fmt.Sprintf("%s %s \n", description, input))
}
