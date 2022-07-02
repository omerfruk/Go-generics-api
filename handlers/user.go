package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/Go-generics-api/model"
	"github.com/omerfruk/Go-generics-api/services"
)

type UserHandler struct {
	BaseHandler[model.User]
	userService services.UserService
}

func NewUserHandler(s services.UserService) UserHandler {
	return UserHandler{
		BaseHandler: BaseHandler[model.User]{
			BaseService: s.BaseService,
		},
		userService: s,
	}
}

func (u *UserHandler) GetByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := u.userService.GetByEmail(email)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(user)
}
