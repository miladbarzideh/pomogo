package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/miladbarzideh/pomogo/configs"
	"github.com/miladbarzideh/pomogo/internal/routes"
)

func main() {

	config, err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	routes.Initialize(app)
	app.Use(logger.New())

	err = app.Listen(fmt.Sprintf(":%v", config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
