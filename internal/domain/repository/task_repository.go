package repository

import (
	"OTakumi/todogo/internal/domain/model"
)

type TaskRepository interface {
	FindAll() ([]*model.Task, error)
	FindByID(id string) (*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(id string) error
}
