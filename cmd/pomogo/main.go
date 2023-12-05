package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/internal/domain/routes"
	"github.com/miladbarzideh/pomogo/internal/infra/config"
	"github.com/miladbarzideh/pomogo/internal/infra/logger"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger.Init(cfg)
	logger.Logger.Info("starting pomogo", zap.String("app-version", cfg.Server.AppVersion))

	app := fiber.New()
	routes.Initialize(app)

	err = app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
