package controllers

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)


func getMedications(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT id, name, time FROM medications")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var medications []gin.H
	for rows.Next() {
		var id int
		var name, time string
		if err := rows.Scan(&id, &name, &time); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		medications = append(medications, gin.H{"id": id, "name": name, "time": time})
	}
	c.JSON(http.StatusOK, medications)
}

func addMedication(db *sql.DB, c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Time string `json:"time"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO medications (name, time) VALUES (?, ?)", req.Name, req.Time)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Medication added"})
}

// Caregiver Support APIs
func getCaregiverStatus(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT id, name, status FROM caregiver_status")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var statuses []gin.H
	for rows.Next() {
		var id int
		var name, status string
		if err := rows.Scan(&id, &name, &status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		statuses = append(statuses, gin.H{"id": id, "name": name, "status": status})
	}
	c.JSON(http.StatusOK, statuses)
}

func updateCaregiverStatus(db *sql.DB, c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO caregiver_status (name, status) VALUES (?, ?)", req.Name, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Status updated"})
}

// System Administration APIs
func getUsers(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT id, name, role FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []gin.H
	for rows.Next() {
		var id int
		var name, role string
		if err := rows.Scan(&id, &name, &role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, gin.H{"id": id, "name": name, "role": role})
	}
	c.JSON(http.StatusOK, users)
}

func addUser(db *sql.DB, c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO users (name, role) VALUES (?, ?)", req.Name, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "User added"})
}

// Emergency Response APIs
func sendEmergencyAlert(c *gin.Context) {
	var req struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Emergency alert sent: %s", req.Message)
	c.JSON(http.StatusOK, gin.H{"status": "Alert sent"})
}