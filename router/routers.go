package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/omerfruk/Go-generics-api/database"
	"github.com/omerfruk/Go-generics-api/handlers"
	"github.com/omerfruk/Go-generics-api/services"
)

func Setup(app fiber.Router) {
	app.Use(logger.New())

	userService := services.NewUserService(database.DB())
	bookService := services.NewBookService(database.DB())

	userHandler := handlers.NewUserHandler(userService)
	bookHandler := handlers.NewBookHandler(bookService)

	app.Get("/users", userHandler.GetAll)
	app.Get("/users/:id", userHandler.GetByID)
	app.Post("/users", userHandler.Create)
	app.Put("/users/:id", userHandler.Update)
	app.Delete("/users/:id", userHandler.Delete)
	app.Get("/users/email/:email", userHandler.GetByEmail)

	app.Get("/books", bookHandler.GetAll)
	app.Get("/books/:id", bookHandler.GetByID)
	app.Post("/books", bookHandler.Create)
	app.Put("/books/:id", bookHandler.Update)
	app.Delete("/books/:id", bookHandler.Delete)
	app.Get("/books/author/:author", bookHandler.GetBookByAuthor)
}
