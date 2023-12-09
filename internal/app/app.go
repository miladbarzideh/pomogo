package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/internal/infra/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	logger *zap.Logger
	cfg    *config.Config
	db     *gorm.DB
}

func NewServer(cfg *config.Config, logger *zap.Logger, db *gorm.DB) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (s *Server) Run() error {
	app := fiber.New()
	s.mapHandlers(app)
	return app.Listen(fmt.Sprintf(":%s", s.cfg.Server.Port))
}
