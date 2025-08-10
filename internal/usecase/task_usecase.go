package usecase

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/domain/service"
	"OTakumi/todogo/internal/repository"
)

type TaskUsecase interface {
	CreateTask(title string) (*model.Task, error)
}

type taskUsecase struct {
	taskRepo    repository.TaskRepository
	idGenerator service.IDGenerator
}

func NewTaskUsecase(tr repository.TaskRepository, ig service.IDGenerator) TaskUsecase {
	return &taskUsecase{
		taskRepo:    tr,
		idGenerator: ig,
	}
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
