package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/internal/domain/project"
)

func MapProjectRoutes(projectGroup fiber.Router, handler project.Handler) {
	projectGroup.Post("", handler.Create())
	projectGroup.Get("/:id", handler.GetByID())
	projectGroup.Get("", handler.GetAll())
	projectGroup.Put("", handler.Update())
	projectGroup.Delete("/:id", handler.DeleteById())
}
