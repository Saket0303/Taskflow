package repository

import (
	"context"
	"taskflow/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	DB *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

// Create project
func (r *ProjectRepository) CreateProject(ctx context.Context, p *models.Project) error {
	query := `
		INSERT INTO projects (name, description, owner_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return r.DB.QueryRow(ctx, query,
		p.Name,
		p.Description,
		p.OwnerID,
	).Scan(&p.ID, &p.CreatedAt)
}

// Get projects for user
func (r *ProjectRepository) GetProjectsByUser(ctx context.Context, userID string) ([]models.Project, error) {
	query := `
		SELECT DISTINCT p.id, p.name, p.description, p.owner_id, p.created_at
		FROM projects p
		LEFT JOIN tasks t ON t.project_id = p.id
		WHERE p.owner_id = $1 OR t.assignee_id = $1
	`

	rows, err := r.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project

	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *ProjectRepository) GetProjectByID(ctx context.Context, id string) (*models.Project, []models.Task, error) {
	projectQuery := `
		SELECT id, name, description, owner_id, created_at
		FROM projects
		WHERE id = $1
	`

	var p models.Project
	err := r.DB.QueryRow(ctx, projectQuery, id).
		Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID, &p.CreatedAt)

	if err != nil {
		return nil, nil, err
	}

	// Get tasks
	taskQuery := `
		SELECT id, title, description, status, priority, project_id, assignee_id, due_date, created_at, updated_at
		FROM tasks
		WHERE project_id = $1
	`

	rows, err := r.DB.Query(ctx, taskQuery, id)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		err := rows.Scan(
			&t.ID, &t.Title, &t.Description, &t.Status,
			&t.Priority, &t.ProjectID, &t.AssigneeID,
			&t.DueDate, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, nil, err
		}
		tasks = append(tasks, t)
	}

	return &p, tasks, nil
}

func (r *ProjectRepository) UpdateProject(ctx context.Context, id string, name, description *string) error {
	query := `
		UPDATE projects
		SET name = COALESCE($1, name),
		    description = COALESCE($2, description)
		WHERE id = $3
	`

	_, err := r.DB.Exec(ctx, query, name, description, id)
	return err
}

func (r *ProjectRepository) DeleteProject(ctx context.Context, id string) error {
	query := `DELETE FROM projects WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}
