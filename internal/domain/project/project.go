package project

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null; size:100"`
	Description string `gorm:"not null; size:200"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Create(project *Project) (*Project, error)
	GetByID(ID uint) (Project, error)
	Update(project Project) (Project, error)
	DeleteById(ID uint) error
}

type Service interface {
	Create(project *Project) (*Project, error)
	GetByID(ID uint) (Project, error)
	Update(project Project) (Project, error)
	DeleteById(ID uint) error
}

type Handler interface {
	Create() fiber.Handler
	GetByID() fiber.Handler
	Update() fiber.Handler
	DeleteById() fiber.Handler
}
