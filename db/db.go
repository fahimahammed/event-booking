package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Assign to global DB

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil { // Ensure DB connection works
		log.Fatalf("Database connection failed: %v", err)
	}

	createTables()
}

func createTables() {
	if DB == nil { // Prevent nil pointer dereference
		log.Fatal("DB is not initialized")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER NOT NULL
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Error creating events table: %v", err)
	}

	log.Println("Database tables initialized successfully")
}
