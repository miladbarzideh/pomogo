package repository

import (
	"github.com/miladbarzideh/pomogo/internal/domain/project"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) project.Repository {
	return &repo{db: db}
}

func (r repo) Create(project *project.Project) (*project.Project, error) {
	result := r.db.Create(project)
	return project, result.Error
}

func (r repo) GetByID(ID uint) (project.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (r repo) Update(project project.Project) (project.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (r repo) DeleteById(ID uint) error {
	//TODO implement me
	panic("implement me")
}
