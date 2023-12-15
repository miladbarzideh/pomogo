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
	tx := r.db.Create(project)
	return project, tx.Error
}

func (r repo) GetByID(id uint) (*project.Project, error) {
	p := &project.Project{ID: id}
	tx := r.db.First(p)
	return p, tx.Error
}

func (r repo) Update(p *project.Project) (*project.Project, error) {
	proj, err := r.GetByID(p.ID)
	if err != nil {
		return proj, err
	}

	tx := r.db.Model(&proj).Updates(project.Project{Title: p.Title, Description: p.Description})
	return proj, tx.Error
}

func (r repo) DeleteById(id uint) error {
	tx := r.db.Delete(&project.Project{}, id)
	return tx.Error
}
