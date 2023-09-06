package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)
func CompleteLesson(lesson string)  {
	 // Open the SQLite database
	 db, err := sql.Open("sqlite3", "pecklin.db")
	 if err != nil {
		 panic(err)
	 }
	 defer db.Close()
	 fmt.Print(lesson)

	    // Insert lesson into the database
        _, err = db.Exec(`
		INSERT OR IGNORE INTO lessons (lesson) VALUES (?);
        `, lesson)
        if err != nil {
            panic(err)
        }
		fmt.Printf("inserted lesson")
		db.Close()
}

func ReadLesson()  {
	
}