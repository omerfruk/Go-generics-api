package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/Go-generics-api/model"
	"gorm.io/gorm"
)

type BookService struct {
	BaseService[model.Book]
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		BaseService[model.Book]{DB: db},
	}
}
func (b BookService) GetBookByAuthor(c *fiber.Ctx) error {
	author := c.Params("author")
	var book *model.Book
	err := b.DB.Model(book).Where("author = ?", author).Find(&book).Error
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(book)
}
