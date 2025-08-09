package repository

import (
	"OTakumi/todogo/internal/domain/model"

	"github.com/google/uuid"
)

type TaskRepository interface {
	FindAll() ([]*model.Task, error)
	FindById(id uuid.UUID) (*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(id uuid.UUID) error
}
