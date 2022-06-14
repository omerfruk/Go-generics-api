package services

import "github.com/omerfruk/Go-generics-api/database"

func GetAll[V any]() ([]V, error) {
	var result []V
	err := database.DB().Find(&result).Error
	return result, err
}
