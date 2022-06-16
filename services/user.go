package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/Go-generics-api/model"
	"gorm.io/gorm"
)

type UserService struct {
	BaseService[model.User]
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		BaseService[model.User]{DB: db},
	}
}

func (u UserService) GetByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	var kullanici *model.User
	err := u.DB.Model(kullanici).Where("email = ?", email).Find(&kullanici).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(kullanici)
}
