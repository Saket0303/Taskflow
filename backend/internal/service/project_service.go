package service

import (
	"context"
	"taskflow/internal/models"
	"taskflow/internal/repository"
)

type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, p *models.Project) error {
	return s.repo.CreateProject(ctx, p)
}

func (s *ProjectService) GetProjects(ctx context.Context, userID string) ([]models.Project, error) {
	return s.repo.GetProjectsByUser(ctx, userID)
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id string) (*models.Project, []models.Task, error) {
	return s.repo.GetProjectByID(ctx, id)
}

func (s *ProjectService) UpdateProject(ctx context.Context, id string, name, description *string) error {
	return s.repo.UpdateProject(ctx, id, name, description)
}

func (s *ProjectService) DeleteProject(ctx context.Context, id string) error {
	return s.repo.DeleteProject(ctx, id)
}
