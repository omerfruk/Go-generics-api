package services

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

type BaseService[T any] struct {
	DB *gorm.DB
}

func (s BaseService[T]) GetAll(c *fiber.Ctx) error {
	var items []T
	err := s.DB.Find(&items).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(items)
}

func (s BaseService[T]) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}
	var item T
	err = s.DB.First(&item, id).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(item)
}

func (s BaseService[T]) Create(c *fiber.Ctx) error {
	var item T
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	err = s.DB.Create(&item).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(item)
}

func (s BaseService[T]) Update(c *fiber.Ctx) error {
	var item T
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	err = s.DB.Save(&item).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(item)
}

func (s BaseService[T]) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}
	var item T
	err = s.DB.Where("id = ?", id).Find(&item).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	err = s.DB.Delete(&item).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.SendStatus(204)
}
