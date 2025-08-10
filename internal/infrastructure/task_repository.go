package infrastructure

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/repository"
	"context"
	"database/sql"
	"fmt"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) FindAll(ctx context.Context) ([]*model.Task, error) {
	query := "SELECT id, title, deadline, is_complete, created_at, updated_at FROM tasks"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		task := &model.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Deadline,
			&task.IsComplete,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}

	return tasks, nil
}

func (r *taskRepository) FindByID(ctx context.Context, id string) (*model.Task, error) {
	// TODO: 実装予定
	return nil, fmt.Errorf("FindByID not implemented yet")
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	// TODO: 実装予定
	return nil, fmt.Errorf("Create not implemented yet")
}

func (r *taskRepository) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	// TODO: 実装予定
	return nil, fmt.Errorf("Update not implemented yet")
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	// TODO: 実装予定
	return fmt.Errorf("Delete not implemented yet")
}
