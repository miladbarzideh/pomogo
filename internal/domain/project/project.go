package project

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `validate:"required" gorm:"not null; size:100"`
	Description string `validate:"required" gorm:"not null; size:200"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Create(project *Project) (*Project, error)
	GetByID(id uint) (*Project, error)
	GetAll() ([]Project, error)
	Update(project *Project) (*Project, error)
	DeleteById(id uint) error
}

type Service interface {
	Create(project *Project) (*Project, error)
	GetByID(id uint) (*Project, error)
	GetAll() ([]Project, error)
	Update(project *Project) (*Project, error)
	DeleteById(id uint) error
}

type Handler interface {
	Create() fiber.Handler
	GetByID() fiber.Handler
	GetAll() fiber.Handler
	Update() fiber.Handler
	DeleteById() fiber.Handler
}
