package services

import (
	"gorm.io/gorm"
)

type IBaseService[t any] interface {
	GetAll()
	Create(t)
	GetById(int)
	Update(t)
	Delete(t)
}

type BaseService[T any] struct {
	DB *gorm.DB
}

func (s BaseService[T]) GetAll() ([]T, error) {
	var result []T
	err := s.DB.Find(&result).Error
	return result, err
}

func (s BaseService[T]) Create(t T) error {
	err := s.DB.Create(&t).Error
	return err
}

func (s BaseService[T]) GetById(id int) (T, error) {
	var result T
	err := s.DB.Where("id = ?", id).First(&result).Error
	return result, err
}

func (s BaseService[T]) Update(id int, t T) error {
	err := s.DB.Model(&t).Where("id = ?", id).Save(&t).Error
	return err
}

func (s BaseService[T]) Delete(id int) error {
	var result T
	err := s.DB.Where("id = ?", id).Delete(&result).Error
	return err
}
