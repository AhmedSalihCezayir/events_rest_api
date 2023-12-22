package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	
	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	
	createTables()
}

func createTables() {
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        nickname TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        is_admin INTEGER 
    )
    `
    _, err := DB.Exec(createUsersTable)
    if err != nil {
        panic("Could not create users table!")
    }

    createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        date_time DATETIME NOT NULL,
        user_id INTEGER,
        FOREIGN KEY(user_id) REFERENCES users(id)
    )
    `
    _, err = DB.Exec(createEventsTable)
    if err != nil {
        panic("Could not create events table!")
    }

    createEventRegistrationsTable := `
    CREATE TABLE IF NOT EXISTS registrations (
        user_id INTEGER NOT NULL,
        event_id INTEGER NOT NULL,
        PRIMARY KEY(user_id, event_id),
        FOREIGN KEY(user_id) REFERENCES users(id),
        FOREIGN KEY(event_id) REFERENCES events(id)
    )
    `
    _, err = DB.Exec(createEventRegistrationsTable)
    if err != nil {
        panic("Could not create registrations table!")
    }
}
