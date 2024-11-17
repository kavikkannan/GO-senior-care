package routes

import (
	"github.com/gofiber/fiber/v2"
	"controllers"
	

)

func Setup(app *fiber.App) {

	app.GET("/health/medications", controllers.getMedications)
	app.POST("/health/medications",)
	app.GET("/caregiver/status", )
	app.POST("/caregiver/status", )
	app.GET("/admin/users", )
	app.POST("/admin/users", )
	app.POST("/emergency/alerts", )
}

