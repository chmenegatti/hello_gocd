package main

import (
	"hello-gocd/src/database"
	"hello-gocd/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.SetupRoutes(app)
	app.Listen(":5000")
}
