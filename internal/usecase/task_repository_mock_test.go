package usecase_test

import (
	"OTakumi/todogo/internal/domain/model"
	"context"

	"github.com/stretchr/testify/mock"
)

// TaskRepositoryのモック
type MockTaskRepository struct {
	mock.Mock
}

// モックがTaskRepositoryインターフェースを実装するように、全てのメソッドを定義する
func (m *MockTaskRepository) FindAll(ctx context.Context) ([]*model.Task, error) {
	args := m.Called(ctx)
	// 戻り値を設定。args.Get(0)が1番目の戻り値、args.Error(1)が2番目の戻り値(error)
	// 戻り値がnilの可能性がある場合は型アサーションで安全に取得する
	var tasks []*model.Task
	if args.Get(0) != nil {
		tasks = args.Get(0).([]*model.Task)
	}
	return tasks, args.Error(1)
}

func (m *MockTaskRepository) FindByID(ctx context.Context, id string) (*model.Task, error) {
	args := m.Called(ctx, id)
	var task *model.Task
	if args.Get(0) != nil {
		task = args.Get(0).(*model.Task)
	}
	return task, args.Error(1)
}

func (m *MockTaskRepository) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	// m.Called に渡された引数を記録する
	args := m.Called(ctx, task)
	var createdTask *model.Task
	if args.Get(0) != nil {
		createdTask = args.Get(0).(*model.Task)
	}
	return createdTask, args.Error(1)
}

func (m *MockTaskRepository) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	args := m.Called(ctx, task)
	var updatedTask *model.Task
	if args.Get(0) != nil {
		updatedTask = args.Get(0).(*model.Task)
	}
	return updatedTask, args.Error(1)
}

func (m *MockTaskRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
