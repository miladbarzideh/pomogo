package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/miladbarzideh/pomogo/internal/routes"
)

func main() {

	app := fiber.New()

	routes.Initialize(app)

	app.Use(logger.New())

	app.Listen(":3000")
}
