package usecase_test

import (
	"OTakumi/todogo/internal/domain/model"
	"OTakumi/todogo/internal/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestTaskUsecase_FindAll は TaskUsecase の FindAll メソッドのテスト
// FindAll メソッドは登録されているすべてのタスクを取得する責務を持つ
func TestTaskUsecase_FindAll(t *testing.T) {
	// テストケース1: タスクが登録されていない場合、空の配列を返す
	t.Run("空のタスクリストを返す", func(t *testing.T) {
		// Arrange: テストの準備
		// MockTaskRepository のインスタンスを作成
		mockRepo := new(MockTaskRepository)

		// IDGeneratorのモックを作成（FindAllでは使用しないが、Usecaseの初期化に必要）
		mockIDGenerator := &MockIDGenerator{ID: "test-id"}

		// テスト用のコンテキストを作成
		ctx := context.Background()

		// リポジトリが返すべき空のタスクリストを定義
		expectedTasks := []*model.Task{}

		// モックの振る舞いを設定：FindAllが呼ばれたら空の配列を返す
		// On メソッドで期待される呼び出しを定義
		// Return メソッドで返却値を指定
		mockRepo.On("FindAll", ctx).Return(expectedTasks, nil)

		// テスト対象のTaskUsecaseインスタンスを作成
		taskUsecase := usecase.NewTaskUsecase(mockRepo, mockIDGenerator)

		// Act: テスト対象のメソッドを実行
		tasks, err := taskUsecase.FindAll(ctx)

		// Assert: 実行結果の検証
		// エラーが発生していないことを確認
		assert.NoError(t, err)

		// 返されたタスクリストが空であることを確認
		assert.Empty(t, tasks)

		// モックに設定した期待値通りに呼び出されたことを確認
		// これにより、リポジトリのFindAllメソッドが確実に呼ばれたことを保証
		mockRepo.AssertExpectations(t)
	})

	// テストケース2: 複数のタスクが登録されている場合、すべてのタスクを返す
	t.Run("複数のタスクを返す", func(t *testing.T) {
		// Arrange: テストの準備
		mockRepo := new(MockTaskRepository)
		mockIDGenerator := &MockIDGenerator{ID: "test-id"}

		ctx := context.Background()

		// テスト用の現在時刻を固定（テストの再現性を保つため）
		now := time.Now()

		// リポジトリが返すべき複数のタスクを定義
		// 異なる状態のタスクを用意することで、様々なケースをカバー
		deadline1 := now
		deadline2 := now.Add(24 * time.Hour)
		expectedTasks := []*model.Task{
			{
				ID:         "1",
				Title:      "Task 1",
				Deadline:   &deadline1,
				IsComplete: false, // 未完了のタスク
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			{
				ID:         "2",
				Title:      "Task 2",
				Deadline:   &deadline2,
				IsComplete: true, // 完了済みのタスク
				CreatedAt:  now,
				UpdatedAt:  now,
			},
		}

		// モックの振る舞いを設定：FindAllが呼ばれたら2件のタスクを返す
		mockRepo.On("FindAll", ctx).Return(expectedTasks, nil)

		taskUsecase := usecase.NewTaskUsecase(mockRepo, mockIDGenerator)

		// Act: テスト対象のメソッドを実行
		tasks, err := taskUsecase.FindAll(ctx)

		// Assert: 実行結果の検証
		// エラーが発生していないことを確認
		assert.NoError(t, err)

		// 返されたタスクの件数が期待値と一致することを確認
		assert.Len(t, tasks, 2)

		// 返されたタスクの内容が期待値と完全に一致することを確認
		// これにより、データが正しく伝播されていることを保証
		assert.Equal(t, expectedTasks, tasks)

		// モックの期待値検証
		mockRepo.AssertExpectations(t)
	})

	// テストケース3: リポジトリ層でエラーが発生した場合、エラーを返す
	t.Run("リポジトリがエラーを返す", func(t *testing.T) {
		// Arrange: テストの準備
		mockRepo := new(MockTaskRepository)
		mockIDGenerator := &MockIDGenerator{ID: "test-id"}

		ctx := context.Background()

		// データベース接続エラーなど、リポジトリ層で発生しうるエラーを定義
		expectedError := errors.New("database error")

		// モックの振る舞いを設定：FindAllが呼ばれたらエラーを返す
		// 第1引数にnil、第2引数にエラーを指定
		mockRepo.On("FindAll", ctx).Return(nil, expectedError)

		taskUsecase := usecase.NewTaskUsecase(mockRepo, mockIDGenerator)

		// Act: テスト対象のメソッドを実行
		tasks, err := taskUsecase.FindAll(ctx)

		// Assert: 実行結果の検証
		// エラーが発生していることを確認
		assert.Error(t, err)

		// 返されたエラーが期待値と一致することを確認
		// これにより、エラーが正しく伝播されていることを保証
		assert.Equal(t, expectedError, err)

		// エラー時はタスクリストがnilであることを確認
		assert.Nil(t, tasks)

		// モックの期待値検証
		mockRepo.AssertExpectations(t)
	})
}
