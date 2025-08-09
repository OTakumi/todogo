package usecase

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/domain/repository"
)

type TaskUsecase interface {
	CreateTask(title string) (*model.Task, error)
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: tr}
}

func (tu *taskUsecase) CreateTask(title string) (*model.Task, error) {
	task := &model.Task{
		Title:      title,
		IsComplete: false,
	}

	return tu.taskRepo.Create(task)
}
