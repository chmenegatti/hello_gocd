package routes

import (
	"hello-gocd/src/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/names")
	api.Get("/", controller.GetAllUsers)
	api.Post("/", controller.CreateUser)
	// api.Patch("/")
	// api.Delete("/")
}
