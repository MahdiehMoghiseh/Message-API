package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/command/:id", controllers.Get_message)

	app.Post("/command", controllers.Create_message)

	app.Put("/command/:id", controllers.Update_message)

	app.Delete("/command/:id", controllers.Delete_message)
}
