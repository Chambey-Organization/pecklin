package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeDatabase() {
	db, err := sql.Open("sqlite3", "pecklin.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS lessons (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            lesson TEXT UNIQUE,
            currentSpeed DOUBLE,
            bestSpeed DOUBLE DEFAULT 0.0
        )
    `)
	if err != nil {
		panic(err)
	}
	db.Close()
}
