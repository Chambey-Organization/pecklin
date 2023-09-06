package database

import (
	 "database/sql"
    _ "github.com/mattn/go-sqlite3"
)


func InitializeDatabase()  {
	db, err := sql.Open("sqlite3", "pecklin.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create the bookings table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS lessons (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            lesson TEXT UNIQUE
        )
    `)
    if err != nil {
        panic(err)
    }
    db.Close()
}