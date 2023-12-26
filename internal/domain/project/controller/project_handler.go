package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/internal/domain/project"
	"github.com/miladbarzideh/pomogo/internal/domain/util"
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
		proj := new(project.Project)
		if err := ctx.BodyParser(proj); err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		result, err := h.projectService.Create(proj)
		if err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (h handler) GetByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := util.GetNumericIdentifier(ctx, "id")
		if err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		result, err := h.projectService.GetByID(uint(id))
		if err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (h handler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		proj := new(project.Project)
		if err := ctx.BodyParser(proj); err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		result, err := h.projectService.Update(proj)
		if err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (h handler) DeleteById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := util.GetNumericIdentifier(ctx, "id")
		if err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		err = h.projectService.DeleteById(uint(id))
		if err != nil {
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusNoContent).JSON("")
	}
}
