package service

import (
	"context"
	"taskflow/internal/models"
	"taskflow/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, t *models.Task) error {
	return s.repo.CreateTask(ctx, t)
}

func (s *TaskService) GetTasks(ctx context.Context, projectID, status, assignee string) ([]models.Task, error) {
	return s.repo.GetTasks(ctx, projectID, status, assignee)
}

func (s *TaskService) UpdateTask(ctx context.Context, id string, t *models.UpdateTaskRequest) error {
	return s.repo.UpdateTask(ctx, id, t)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.DeleteTask(ctx, id)
}

func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}
