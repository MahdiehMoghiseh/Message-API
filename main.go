package main

import (
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.UserRoute(app)

	app.Listen(":1234")
}
