package model

import (
	"github.com/omerfruk/Go-generics-api/services"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService struct {
	services.BaseService[User]
}

func (s UserService) GetAll() ([]User, error) {
	users, err := s.GetAll()
	return users, err
}
