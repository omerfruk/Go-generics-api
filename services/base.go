package services

import (
	"gorm.io/gorm"
)

type IBaseService interface {
	GetAll() ([]interface{}, error)
	GetById(id int) (interface{}, error)
	Create(item interface{}) (interface{}, error)
	Update(item interface{}) (interface{}, error)
	Delete(item interface{}) error
}

type BaseService[T any] struct {
	DB *gorm.DB
}

func (s BaseService[T]) GetAll() ([]T, error) {
	var items []T
	err := s.DB.Find(&items).Error
	return items, err
}

func (s BaseService[T]) GetById(id int) (T, error) {
	var item T
	err := s.DB.First(&item, id).Error
	return item, err
}

func (s BaseService[T]) Create(item T) (T, error) {
	err := s.DB.Create(&item).Error
	return item, err
}

func (s BaseService[T]) Update(item T) (T, error) {
	err := s.DB.Save(&item).Error
	return item, err
}

func (s BaseService[T]) Delete(item T) error {
	err := s.DB.Delete(&item).Error
	return err
}
