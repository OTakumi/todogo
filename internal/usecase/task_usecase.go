package usecase

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/domain/repository"
	"OTakumi/todogo/internal/domain/service"
)

type TaskUsecase interface {
	CreateTask(title string) (*model.Task, error)
}

type taskUsecase struct {
	taskRepo    repository.TaskRepository
	idGenerator service.IDGenerator
}

func NewTaskUsecase(tr repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: tr}
}

func (tu *taskUsecase) CreateTask(title string) (*model.Task, error) {
	// idを取得する
	id := tu.idGenerator.NewID()

	// タスクを生成
	task := model.NewTask(id, title)

	if err := task.Validate(); err != nil {
		return nil, err
	}

	return tu.taskRepo.Create(task)
}
