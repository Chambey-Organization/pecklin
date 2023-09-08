package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"main.go/pkg/models"
)
func CompleteLesson(lesson models.LessonDTO)  {
	 // Open the SQLite database
	 db, err := sql.Open("sqlite3", "pecklin.db")
	 if err != nil {
		 panic(err)
	 }
	 defer db.Close()

	    // Insert lesson into the database
        _, err = db.Exec(`
		INSERT OR IGNORE INTO lessons (lesson, speed) VALUES (?,?);
        `, lesson.Title, lesson.Speed)
        if err != nil {
            panic(err)
        }
		db.Close()
}

func ReadCompletedLesson()[]models.LessonDTO  {
	 db, err := sql.Open("sqlite3", "pecklin.db")
	 if err != nil {
		 panic(err)
	 }
	 defer db.Close()

	 rows, err := db.Query("SELECT lesson, speed FROM lessons")
	 if err != nil {
		 panic(err)
	 }
	 defer rows.Close()
	 var lesson models.LessonDTO
	 var lessons []models.LessonDTO

	 for rows.Next() {		
		 err := rows.Scan(&lesson.Title, &lesson.Speed)
		 if err != nil {
			 return lessons
		 }
		 lessons = append(lessons, lesson)
	 }
		db.Close()
		return lessons
}