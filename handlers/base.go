package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/Go-generics-api/services"
)

type BaseHandler[T any] struct {
	BaseService services.BaseService[T]
}

func (b *BaseHandler[T]) GetAll(c *fiber.Ctx) error {
	data, err := b.BaseService.GetAll()
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(data)
}

func (b *BaseHandler[T]) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(err)
	}
	data, err := b.BaseService.GetById(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(data)
}

func (b *BaseHandler[T]) Create(c *fiber.Ctx) error {
	var item T
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	data, err := b.BaseService.Create(item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(data)
}

func (b *BaseHandler[T]) Update(c *fiber.Ctx) error {
	var item T
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	data, err := b.BaseService.Update(item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(data)
}

func (b *BaseHandler[T]) Delete(c *fiber.Ctx) error {
	var item T
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	err = b.BaseService.Delete(item)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(nil)
}
