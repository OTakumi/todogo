package infrastructure

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/repository"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	defer func() {
		if err := rows.Close(); err != nil {
			// Log the error if needed, but don't override the main error
			_ = err
		}
	}()

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
	// コンテキストの確認
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("context cancelled before task creation: %w", ctx.Err())
	default:
	}

	// バリデーション
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// タスクのコピーを作成（元のオブジェクトを変更しないため）
	newTask := *task

	// IDの処理
	if newTask.ID == "" {
		newTask.ID = uuid.New().String()
	} else {
		// 既存のIDがある場合、UUID形式であることを検証
		if _, err := uuid.Parse(newTask.ID); err != nil {
			return nil, fmt.Errorf("invalid task ID format: %w", err)
		}
	}

	// タイムスタンプの設定
	now := time.Now()
	newTask.CreatedAt = now
	newTask.UpdatedAt = now

	// トランザクションを開始
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// SQLクエリの実行
	query := `
		INSERT INTO tasks (id, title, deadline, is_complete, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.ExecContext(ctx, query,
		newTask.ID,
		newTask.Title,
		newTask.Deadline,
		newTask.IsComplete,
		newTask.CreatedAt,
		newTask.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert task: %w", err)
	}

	// トランザクションのコミット
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &newTask, nil
}

func (r *taskRepository) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	// TODO: 実装予定
	return nil, fmt.Errorf("Update not implemented yet")
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	// TODO: 実装予定
	return fmt.Errorf("Delete not implemented yet")
}
