package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"main.go/pkg/models"
)

func CompleteLesson(lesson models.LessonDTO) {
	db, err := sql.Open("sqlite3", "pecklin.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert the lesson if it doesn't exist, or ignore if it does.
	_, err = db.Exec(`
        INSERT OR REPLACE INTO lessons (lesson, currentSpeed) VALUES (?, ?);
    `, lesson.Title, lesson.CurrentSpeed)
	if err != nil {
		panic(err.Error())
	}

	lessonToUpdate := ReadALessonData(lesson)

	if lesson.CurrentSpeed > lessonToUpdate.BestSpeed {
		fmt.Printf("lessons -> %s %.2F %.2F \n", lesson.Title, lesson.CurrentSpeed, lesson.BestSpeed)

		_, err = db.Exec(
			`UPDATE lessons SET bestSpeed = ? WHERE  lesson = ?;
    `, lesson.CurrentSpeed, lesson.Title)
	}
	db.Close()
}

func ReadCompletedLesson() []models.LessonDTO {
	db, err := sql.Open("sqlite3", "pecklin.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT lesson, currentSpeed FROM lessons")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var lesson models.LessonDTO
	var lessons []models.LessonDTO

	for rows.Next() {
		err := rows.Scan(&lesson.Title, &lesson.CurrentSpeed)

		if err != nil {
			return lessons
		}
		lessons = append(lessons, lesson)
	}
	db.Close()
	return lessons
}

func ReadALessonData(searchLesson models.LessonDTO) models.LessonDTO {
	db, err := sql.Open("sqlite3", "pecklin.db")
	query := "SELECT lesson, currentSpeed, bestSpeed FROM lessons WHERE lesson = ?"

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(query, searchLesson.Title)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var lesson models.LessonDTO

	for rows.Next() {
		err := rows.Scan(&lesson.Title, &lesson.CurrentSpeed, &lesson.BestSpeed)
		if err != nil {
			return lesson
		}
	}
	db.Close()
	return lesson
}
