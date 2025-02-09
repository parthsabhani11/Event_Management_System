package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	fmt.Print("Connecting to database...\n")
	DB, err = sql.Open("sqlite3", "eventmanager.db")

	if err != nil {
		panic("Unable to access database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	fmt.Print("Connection to database success!\n")

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'user'
	)`

	fmt.Print("Creating User table...")
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Unable to create User table")
	}
	fmt.Print("Success!\n")

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
		FOREIGN KEY(userId) REFERENCES users(id)
	)`

	fmt.Print("Creating Event table...")
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		fmt.Print("Failed!\n")
		panic("Unable to create Event table")
	}
	fmt.Print("Success!\n")

	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	fmt.Print("Creating Registration table...")
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		fmt.Print("Failed!\n")
		panic(err)
	}
	fmt.Print("Success!\n")
}
