package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/omerfruk/Go-generics-api/database"
	"github.com/omerfruk/Go-generics-api/services"
)

func Setup(app fiber.Router) {
	app.Use(logger.New())

	userService := services.NewUserService(database.DB())
	bookService := services.NewBookService(database.DB())

	app.Get("/users", userService.GetAll)
	app.Get("/users/:id", userService.GetById)
	app.Post("/users", userService.Create)
	// Todo: Burada id almamıza gerek yok
	app.Post("/users/:id", userService.Update)
	app.Delete("/users/:id", userService.Delete)
	app.Get("/users/email/:email", userService.GetByEmail)

	app.Get("/books", bookService.GetAll)
	app.Get("/books/:id", bookService.GetById)
	app.Post("/books", bookService.Create)
	// Todo: Burada id almamıza gerek yok
	app.Post("/books/:id", bookService.Update)
	app.Delete("/books/:id", bookService.Delete)
	app.Get("/books/author/:author", bookService.GetBookByAuthor)
}
