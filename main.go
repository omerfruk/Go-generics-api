package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerfruk/Go-generics-api/database"
	"github.com/omerfruk/Go-generics-api/router"
)

func main() {
	//DB connection and auto migrate area
	database.DBConnect()
	database.AutoMigrate()

	//Router area
	app := fiber.New()
	router.Setup(app)
	app.Listen(":4747")
}
