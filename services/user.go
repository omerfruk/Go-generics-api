package services

import (
	"github.com/omerfruk/Go-generics-api/model"
	"gorm.io/gorm"
)

type UserService struct {
	BaseService[model.User]
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{
		BaseService[model.User]{DB: db},
	}
}

func (u UserService) GetByEmail(email string) (model.User, error) {
	var kullanici model.User
	err := u.DB.Where("email = ?", email).Find(&kullanici).Error
	return kullanici, err
}
