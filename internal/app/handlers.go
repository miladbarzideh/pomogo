package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/miladbarzideh/pomogo/api"
	"github.com/miladbarzideh/pomogo/internal/domain/project/controller"
	"github.com/miladbarzideh/pomogo/internal/domain/project/repository"
	"github.com/miladbarzideh/pomogo/internal/domain/project/service"
)

func (s *Server) mapHandlers(app *fiber.App) error {

	// Init repositories
	projectRepository := repository.NewProjectRepository(s.db)

	// Init services (use cases)
	projectService := service.NewProjectService(projectRepository, s.logger)

	// Init handlers
	projectHandler := controller.NewProjectHandler(projectService, s.logger)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Map routes
	v1 := app.Group("/api/v1")
	projectGroup := v1.Group("/projects")
	controller.MapProjectRoutes(projectGroup, projectHandler)

	return nil
}
