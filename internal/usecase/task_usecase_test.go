package usecase_test

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// タスク作成が成功する場合
func TestTaskUsecase_CreateTask_Failure(t *testing.T) {
	// Arrange
	mockRepo := new(MockTaskRepository)

	// 固定のIDを返すモックIDジェネレータ
	expectedUUID := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	expectedTitle := "usecaseのテストを書く"

	mockIDGenerator := &MockIDGenerator{ID: expectedUUID}

	mockRepo.On(
		"Create",
		mock.MatchedBy(func(task *model.Task) bool {
			idMatch := task.ID == expectedUUID
			titleMatch := task.Title == expectedTitle
			return idMatch && titleMatch
		}),
	).Return(nil, errors.New("error"))

	// UsecaseにRepositoryとIDGeneratorのモックを注入
	taskUsecase := usecase.NewTaskUsecase(mockRepo, mockIDGenerator)

	// Act
	_, err := taskUsecase.CreateTask(expectedTitle)

	// Assert
	// エラーがないことを確認
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}

// リポジトリがエラーを返し、タスク作成が失敗する場合
