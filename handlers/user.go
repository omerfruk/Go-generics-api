package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/Go-generics-api/model"
)

func GetAll(c *fiber.Ctx) error {
	var user model.UserService
	users, err := user.GetAll()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(users)
}
