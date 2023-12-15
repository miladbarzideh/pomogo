package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/internal/domain/project"
	"go.uber.org/zap"
)

type handler struct {
	projectService project.Service
	logger         *zap.Logger
}

func NewProjectHandler(projectService project.Service, logger *zap.Logger) project.Handler {
	return &handler{
		projectService: projectService,
		logger:         logger,
	}
}

func (h handler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(project.Project)
		err := ctx.BodyParser(request)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}
		result, err := h.projectService.Create(request)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": result}})
	}
}

func (h handler) GetByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Implement me")
	}
}

func (h handler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Implement me")
	}
}

func (h handler) DeleteById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Implement me")
	}
}
