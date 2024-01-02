package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/miladbarzideh/pomogo/internal/domain/project"
	"github.com/miladbarzideh/pomogo/internal/domain/util"
	"go.uber.org/zap"
)

type handler struct {
	projectService project.Service
	logger         *zap.Logger
}

var validate = validator.New()

func NewProjectHandler(projectService project.Service, logger *zap.Logger) project.Handler {
	return &handler{
		projectService: projectService,
		logger:         logger,
	}
}

// Create project godoc
// @Summary      Create a project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  project.Project
// @Failure      422  {object}  util.ResponseError
// @Failure      400  {object}  util.ResponseError
// @Router       /projects [post]
func (h handler) Create() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		proj := new(project.Project)
		if err := ctx.BodyParser(proj); err != nil {
			h.logger.Error("failed to parse body", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		if err := validate.Struct(proj); err != nil {
			h.logger.Error("validation errors", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		result, err := h.projectService.Create(proj)
		if err != nil {
			h.logger.Error("failed to create the project", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (h handler) GetByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := util.GetNumericIdentifier(ctx, "id")
		if err != nil {
			h.logger.Error("failed to parse numeric ID", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		result, err := h.projectService.GetByID(uint(id))
		if err != nil {
			h.logger.Error("failed to find the project", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (h handler) GetAll() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		projects, err := h.projectService.GetAll()
		if err != nil {
			h.logger.Error("failed to get all projects", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusOK).JSON(projects)
	}
}

func (h handler) Update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		proj := new(project.Project)
		if err := ctx.BodyParser(proj); err != nil {
			h.logger.Error("failed to parse body", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		if err := validate.Struct(proj); err != nil {
			h.logger.Error("validation errors", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		result, err := h.projectService.Update(proj)
		if err != nil {
			h.logger.Error("failed to update the project", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (h handler) DeleteById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := util.GetNumericIdentifier(ctx, "id")
		if err != nil {
			h.logger.Error("failed to parse numeric ID", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		err = h.projectService.DeleteById(uint(id))
		if err != nil {
			h.logger.Error("failed to delete the project", zap.Error(err))
			return fiber.NewError(util.ParseError(err))
		}

		return ctx.Status(fiber.StatusNoContent).JSON("")
	}
}
