package usecase

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/domain/service"
	"OTakumi/todogo/internal/repository"
	"context"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, title string) (*model.Task, error)
	FindAll(ctx context.Context) ([]*model.Task, error)
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

func (tu *taskUsecase) CreateTask(ctx context.Context, title string) (*model.Task, error) {
	// idを取得する
	id := tu.idGenerator.NewID()

	// タスクを生成
	task := model.NewTask(id, title)

	if err := task.Validate(); err != nil {
		return nil, err
	}

	return tu.taskRepo.Create(ctx, task)
}

// FindAll は登録されているすべてのタスクを取得する
// リポジトリ層に処理を委譲し、取得したタスクをそのまま返す
func (tu *taskUsecase) FindAll(ctx context.Context) ([]*model.Task, error) {
	// リポジトリ層のFindAllメソッドを呼び出し
	// データベースからすべてのタスクを取得する
	return tu.taskRepo.FindAll(ctx)
}
