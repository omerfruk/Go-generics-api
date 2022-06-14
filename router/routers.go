package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app fiber.Router) {
	app.Use(logger.New)

	app.Get("users")
}
