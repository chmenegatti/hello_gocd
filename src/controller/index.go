package controller

import (
	"hello-gocd/src/database"
	"hello-gocd/src/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUsersResponse(userModel models.Names) User {
	return User{Id: userModel.Id, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []models.Names{}

	database.Database.Db.Find(&users)
	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateUsersResponse(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}
