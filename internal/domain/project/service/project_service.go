package service

import (
	"github.com/miladbarzideh/pomogo/internal/domain/project"
	"go.uber.org/zap"
)

type service struct {
	projectRepo project.Repository
	logger      *zap.Logger
}

func NewProjectService(projectRepo project.Repository, logger *zap.Logger) project.Service {
	return &service{
		projectRepo: projectRepo,
		logger:      logger,
	}
}

func (s service) Create(project *project.Project) (*project.Project, error) {
	return s.projectRepo.Create(project)
}

func (s service) GetByID(id uint) (*project.Project, error) {
	return s.projectRepo.GetByID(id)
}

func (s service) Update(project *project.Project) (*project.Project, error) {
	return s.projectRepo.Update(project)
}

func (s service) DeleteById(id uint) error {
	return s.projectRepo.DeleteById(id)
}
