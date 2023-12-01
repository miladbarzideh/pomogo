package logger

import (
	"log"

	"github.com/miladbarzideh/pomogo/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Init Init Logger
func Init(cfg *configs.Config) {
	logLevel := getLoggerLevel(cfg)

	config := zap.NewDevelopmentConfig()

	config.Level.SetLevel(logLevel)
	logger, err := config.Build()
	if err != nil {
		log.Fatal("Could not create Logger")
	}
	Logger = logger
}

func getLoggerLevel(cfg *configs.Config) zapcore.Level {
	var level zapcore.Level
	if err := level.Set(cfg.Server.LogLevel); err != nil {
		log.Printf("Invalid log level %s", cfg.Server.LogLevel)
		level = zapcore.InfoLevel
	}
	return level
}
