package repository

import (
	"OTakumi/todogo/internal/domain/model"
	"context"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]*model.Task, error)
	FindByID(ctx context.Context, id string) (*model.Task, error)
	Create(ctx context.Context, task *model.Task) (*model.Task, error)
	Update(ctx context.Context, task *model.Task) (*model.Task, error)
	Delete(ctx context.Context, id string) error
}
