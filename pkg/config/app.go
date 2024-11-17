package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
	


// Connect opens the SQLite database and creates tables if they don't exist
func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}


	// Create tables if they don't exist
	err = createTables(DB)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

// createTables runs SQL statements to create each table if it doesn't exist

func createTables(db *sql.DB) error {
    tableStatements := []string{
        // User Accounts Table
        `CREATE TABLE IF NOT EXISTS medications (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		time TEXT
	);`,
	`CREATE TABLE IF NOT EXISTS caregiver_status (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		status TEXT
	);`,
	`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		role TEXT
	);`,
    }

    for _, statement := range tableStatements {
        _, err := db.Exec(statement)
        if err != nil {
            return err
        }
    }
    return nil
}


// GetDB returns a pointer to the database connection
func GetDB() *sql.DB {
	return DB
}
