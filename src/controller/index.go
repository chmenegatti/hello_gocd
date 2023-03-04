package controller

import (
	users "hello-gocd/src/services/Users"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) (err error) {

	users := users.GetAllUsersService

	return c.Status(200).JSON(users)
}
