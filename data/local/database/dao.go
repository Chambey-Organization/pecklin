package database

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"main.go/domain/models"
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

func InsertPractices(practices []models.Practice) {
	for _, practice := range practices {
		//var existingPractice models.Practice
		DB.Save(&practice)
		//DB.Where("id = ?", practice.ID).FirstOrCreate(&existingPractice, &practice)
	}
}

func ReadPractices() []models.Practice {
	var practices []models.Practice
	DB.Find(&practices)
	return practices
}

func ReadAllLessons() []models.Lesson {
	var lessons []models.Lesson
	DB.Find(&lessons)

	fmt.Println("------------------------ lessons ----------------------")
	for _, lesson := range lessons {
		fmt.Printf("ID: %d, practiceId: %d, Name: %s\n", lesson.ID, lesson.PracticeID, lesson.Title)
	}
	return lessons
}

func ReadAllLessonsContent() []models.LessonContent {
	var lessons []models.LessonContent
	DB.Find(&lessons)

	fmt.Println("------------------------ lesson content ----------------------")
	for _, lesson := range lessons {
		fmt.Printf("ID: %d, LessonId: %d, Name: %s\n", lesson.ID, lesson.LessonID, lesson.Prompt)
	}
	return lessons
}
