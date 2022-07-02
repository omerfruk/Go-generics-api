package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/Go-generics-api/model"
	"github.com/omerfruk/Go-generics-api/services"
)

type BookHandler struct {
	BaseHandler[model.Book]
	bookService services.BookService
}

func NewBookHandler(s services.BookService) BookHandler {
	return BookHandler{
		BaseHandler: BaseHandler[model.Book]{
			BaseService: s.BaseService,
		},
		bookService: s,
	}
}

func (b *BookHandler) GetBookByAuthor(c *fiber.Ctx) error {
	author := c.Params("author")
	book, err := b.bookService.GetBookByAuthor(author)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.JSON(book)
}
