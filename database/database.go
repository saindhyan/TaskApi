package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

// connection with sqlite database file
func Connect() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	CreateTable()
	return db, err
}

// creating table its not exists
func CreateTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS task (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		dueDate TEXT,
		status TEXT
	);
	`)
	if err != nil {
		log.Fatal("error in creating table :", err)

	}

}
