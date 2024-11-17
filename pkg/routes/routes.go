package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/codeknight05/GO-senior-care/pkg/controllers"

)

func Setup(app *fiber.App) {

	app.GET("/health/medications", controllers.ge)
	app.POST("/health/medications",)
	app.GET("/caregiver/status", )
	app.POST("/caregiver/status", )
	app.GET("/admin/users", )
	app.POST("/admin/users", )
	app.POST("/emergency/alerts", )
}

