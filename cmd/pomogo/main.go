package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/configs"
	"github.com/miladbarzideh/pomogo/internal/logger"
	"github.com/miladbarzideh/pomogo/internal/routes"
	"go.uber.org/zap"
)

func main() {

	config, err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger.Init(config)
	logger.Logger.Info("starting pomogo", zap.String("app-version", config.Server.AppVersion))

	app := fiber.New()
	routes.Initialize(app)

	err = app.Listen(fmt.Sprintf(":%s", config.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
