package repository

import (
	"context"
	"taskflow/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	DB *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{DB: db}
}

// Create Task
func (r *TaskRepository) CreateTask(ctx context.Context, t *models.Task) error {
	query := `
		INSERT INTO tasks (title, description, status, priority, project_id, assignee_id, due_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`

	return r.DB.QueryRow(ctx, query,
		t.Title,
		t.Description,
		t.Status,
		t.Priority,
		t.ProjectID,
		t.AssigneeID,
		t.DueDate,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

// Get Tasks with filters
func (r *TaskRepository) GetTasks(ctx context.Context, projectID, status, assignee string) ([]models.Task, error) {
	query := `
		SELECT id, title, description, status, priority, project_id, assignee_id, due_date, created_at, updated_at
		FROM tasks
		WHERE project_id = $1
	`

	args := []interface{}{projectID}

	if status != "" {
		query += " AND status = $2"
		args = append(args, status)
	}

	if assignee != "" {
		query += " AND assignee_id = $3"
		args = append(args, assignee)
	}

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, id string, t *models.UpdateTaskRequest) error {
	query := `
		UPDATE tasks
		SET title = COALESCE($1, title),
		    description = COALESCE($2, description),
		    status = COALESCE($3, status),
		    priority = COALESCE($4, priority),
		    assignee_id = COALESCE($5, assignee_id),
		    due_date = COALESCE($6, due_date),
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $7
	`

	_, err := r.DB.Exec(ctx, query,
		t.Title,
		t.Description,
		t.Status,
		t.Priority,
		t.AssigneeID,
		t.DueDate,
		id,
	)

	return err
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`

	_, err := r.DB.Exec(ctx, query, id)
	return err
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	query := `
		SELECT id, project_id
		FROM tasks
		WHERE id = $1
	`

	var t models.Task
	err := r.DB.QueryRow(ctx, query, id).
		Scan(&t.ID, &t.ProjectID)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
