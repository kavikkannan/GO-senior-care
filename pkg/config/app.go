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
        `CREATE TABLE IF NOT EXISTS Users (
            user_id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            password_hash TEXT NOT NULL,
            role TEXT CHECK(role IN ('senior', 'caregiver', 'healthcare_provider', 'admin')) NOT NULL,
            date_created DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,

        // Health Monitoring: Medication Management
        `CREATE TABLE IF NOT EXISTS Medications (
            medication_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            medication_name TEXT NOT NULL,
            dosage TEXT,
            frequency TEXT,
            start_date DATE,
            end_date DATE,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // Health Monitoring: Reminders
        `CREATE TABLE IF NOT EXISTS Reminders (
            reminder_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            medication_id INTEGER,
            reminder_time TIME NOT NULL,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE,
            FOREIGN KEY (medication_id) REFERENCES Medications(medication_id) ON DELETE SET NULL
        );`,

        // Health Monitoring: Sleep Pattern Monitoring
        `CREATE TABLE IF NOT EXISTS SleepPatterns (
            sleep_pattern_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            sleep_start DATETIME,
            sleep_end DATETIME,
            duration INTEGER,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // Caregiver Support: Status Updates to Caregivers
        `CREATE TABLE IF NOT EXISTS StatusUpdates (
            status_update_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            caregiver_id INTEGER NOT NULL,
            update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
            status_message TEXT,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE,
            FOREIGN KEY (caregiver_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // Caregiver Support: Communication Platform
        `CREATE TABLE IF NOT EXISTS Messages (
            message_id INTEGER PRIMARY KEY AUTOINCREMENT,
            sender_id INTEGER NOT NULL,
            receiver_id INTEGER NOT NULL,
            message_text TEXT,
            sent_time DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (sender_id) REFERENCES Users(user_id) ON DELETE CASCADE,
            FOREIGN KEY (receiver_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // System Administration: Performance Monitoring
        `CREATE TABLE IF NOT EXISTS PerformanceMetrics (
            metric_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            metric_type TEXT,
            metric_value REAL,
            recorded_time DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // Emergency Response: Emergency Alerts
        `CREATE TABLE IF NOT EXISTS EmergencyAlerts (
            alert_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            alert_type TEXT CHECK(alert_type IN ('fall', 'health_issue', 'other')) NOT NULL,
            alert_time DATETIME DEFAULT CURRENT_TIMESTAMP,
            resolved BOOLEAN DEFAULT 0,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // Emergency Response: Fall Detection
        `CREATE TABLE IF NOT EXISTS FallDetections (
            fall_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            fall_time DATETIME DEFAULT CURRENT_TIMESTAMP,
            location TEXT,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // System Administration: User Account Management Audit Log
        `CREATE TABLE IF NOT EXISTS AuditLog (
            log_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            action TEXT,
            action_time DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
        );`,

        // System Administration: User Session Table for User Management and Monitoring
        `CREATE TABLE IF NOT EXISTS UserSessions (
            session_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            login_time DATETIME DEFAULT CURRENT_TIMESTAMP,
            logout_time DATETIME,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
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
