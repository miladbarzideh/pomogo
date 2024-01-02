package main

import (
	"log"

	"github.com/miladbarzideh/pomogo/internal/app"
	"github.com/miladbarzideh/pomogo/internal/infra/config"
	"github.com/miladbarzideh/pomogo/internal/infra/db"
	"github.com/miladbarzideh/pomogo/internal/infra/logger"
	"go.uber.org/zap"
)

// @title Pomogo
// @version 1.0
// @description A To-Do list based on pomodoro technique

// @contact.name Milad Barzideh
// @contact.email barzidehmilad@gmail.com
// @contact.url https://github.com/miladbarzideh

// @host      localhost:8513
// @BasePath /api/v1
func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	zLog := logger.InitLogger(cfg)
	zLog.Info("starting pomogo", zap.String("app-version", cfg.Server.AppVersion))

	postgresDb, err := db.NewConnection(cfg)
	if err != nil {
		zLog.Fatal("database connection failed", zap.Error(err))
	}

	server := app.NewServer(cfg, zLog, postgresDb)
	if server.Run() != nil {
		zLog.Fatal("failed to start the app", zap.Error(err))
	}
}
