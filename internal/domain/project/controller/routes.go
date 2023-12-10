package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/internal/domain/project"
)

func MapProjectRoutes(projectGroup fiber.Router, handler project.Handler) {
	projectGroup.Post("", handler.Create())
	projectGroup.Get("", handler.GetByID())
	projectGroup.Post("", handler.Update())
	projectGroup.Post("", handler.DeleteById())
}
