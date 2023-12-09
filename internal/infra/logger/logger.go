package logger

import (
	"log"

	"github.com/miladbarzideh/pomogo/internal/infra/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger Init Logger
func InitLogger(cfg *config.Config) *zap.Logger {
	logLevel := getLoggerLevel(cfg)
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.Level.SetLevel(logLevel)
	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatal("Could not create Logger")
	}
	return logger
}

func getLoggerLevel(cfg *config.Config) zapcore.Level {
	var level zapcore.Level
	if err := level.Set(cfg.Server.LogLevel); err != nil {
		log.Printf("Invalid log level %s", cfg.Server.LogLevel)
		level = zapcore.InfoLevel
	}
	return level
}
