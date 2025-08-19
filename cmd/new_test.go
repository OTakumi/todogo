package cmd

import (
	"OTakumi/todogo/internal/domain/model"
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTaskUsecase はTaskUsecaseインターフェースのモック実装
// testifyのmock.Mockを埋め込んで、メソッド呼び出しの記録と検証を可能にする
type MockTaskUsecase struct {
	mock.Mock
}

// CreateTask はTaskUsecaseインターフェースのCreateTaskメソッドのモック実装
// 引数: ctx - コンテキスト, title - タスクのタイトル
// 戻り値: 作成されたタスク, エラー
func (m *MockTaskUsecase) CreateTask(ctx context.Context, title string) (*model.Task, error) {
	// Calledメソッドで呼び出しを記録し、事前に設定された戻り値を取得
	args := m.Called(ctx, title)

	// 最初の戻り値がnilの場合、nilとエラーを返す
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	// 正常な場合、タスクとエラー（通常はnil）を返す
	return args.Get(0).(*model.Task), args.Error(1)
}

// FindAll はTaskUsecaseインターフェースのFindAllメソッドのモック実装
// このテストでは使用されないが、インターフェースを満たすために必要
func (m *MockTaskUsecase) FindAll(ctx context.Context) ([]*model.Task, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Task), args.Error(1)
}

// TestNewCommand_CreateTaskWithTitle は正常にタスクを作成できることを確認するテスト
// タイトルを指定してnewコマンドを実行し、期待通りの動作をすることを検証
func TestNewCommand_CreateTaskWithTitle(t *testing.T) {
	// Arrange: テストの準備
	// モックのUsecaseを作成
	mockUsecase := new(MockTaskUsecase)

	// 元のtaskUsecaseを保存し、テスト用のモックに置き換える
	// defer文でテスト終了時に元に戻すことで、他のテストに影響を与えない
	originalTaskUsecase := taskUsecase
	taskUsecase = mockUsecase
	defer func() { taskUsecase = originalTaskUsecase }()

	// CreateTaskメソッドが呼ばれた時の戻り値を設定
	// "Test Task"というタイトルで呼ばれたら、作成されたタスクを返すように設定
	createdTask := &model.Task{
		ID:    "test-id-123",
		Title: "Test Task",
	}
	mockUsecase.On("CreateTask", mock.Anything, "Test Task").Return(createdTask, nil)

	// コマンドの出力をキャプチャするためのバッファを作成
	// 標準出力と標準エラー出力の両方をこのバッファにリダイレクト
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Act: テスト対象の実行
	// newコマンドに--titleフラグを付けて実行
	rootCmd.SetArgs([]string{"new", "--title", "Test Task"})
	err := rootCmd.Execute()

	// Assert: 結果の検証
	// エラーが発生していないことを確認
	assert.NoError(t, err)

	// 出力に成功メッセージが含まれていることを確認
	assert.Contains(t, buf.String(), "Task created successfully")

	// 作成されたタスクのIDが出力に含まれていることを確認
	assert.Contains(t, buf.String(), "test-id-123")

	// モックが期待通りに呼び出されたことを検証
	mockUsecase.AssertExpectations(t)
}

// TestNewCommand_ErrorWhenTitleIsEmpty は空のタイトルを拒否することを確認するテスト
// バリデーションが正しく機能していることを検証
func TestNewCommand_ErrorWhenTitleIsEmpty(t *testing.T) {
	// Arrange: テストの準備
	// 出力キャプチャ用のバッファを設定
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Act: テスト対象の実行
	// 空文字列のタイトルでnewコマンドを実行
	rootCmd.SetArgs([]string{"new", "--title", ""})
	err := rootCmd.Execute()

	// Assert: 結果の検証
	// エラーが発生することを確認
	assert.Error(t, err)

	// エラーメッセージに適切な内容が含まれていることを確認
	assert.Contains(t, buf.String(), "title cannot be empty")
}

// TestNewCommand_ErrorWhenNoTitleProvided はタイトルフラグが必須であることを確認するテスト
// 必須フラグが設定されていない場合のエラーハンドリングを検証
func TestNewCommand_ErrorWhenNoTitleProvided(t *testing.T) {
	// Arrange: テストの準備
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Act: テスト対象の実行
	// titleフラグなしでnewコマンドを実行
	rootCmd.SetArgs([]string{"new"})
	err := rootCmd.Execute()

	// Assert: 結果の検証
	// エラーが発生することを確認
	assert.Error(t, err)

	// タイトルが空の場合のエラーメッセージを確認（Cobraの動作により変更）
	// フラグが設定されていない場合も、空文字列として扱われるため同じエラーメッセージになる
	assert.Contains(t, buf.String(), "title cannot be empty")
}

// TestNewCommand_HandleUsecaseError はUsecaseレイヤーでエラーが発生した場合の処理を確認するテスト
// データベースエラーなど、ビジネスロジック層のエラーが適切に処理されることを検証
func TestNewCommand_HandleUsecaseError(t *testing.T) {
	// Arrange: テストの準備
	mockUsecase := new(MockTaskUsecase)
	originalTaskUsecase := taskUsecase
	taskUsecase = mockUsecase
	defer func() { taskUsecase = originalTaskUsecase }()

	// CreateTaskメソッドがエラーを返すように設定
	// assert.AnErrorは汎用的なエラーオブジェクト
	mockUsecase.On("CreateTask", mock.Anything, "Test Task").Return(nil, assert.AnError)

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Act: テスト対象の実行
	rootCmd.SetArgs([]string{"new", "--title", "Test Task"})
	err := rootCmd.Execute()

	// Assert: 結果の検証
	// エラーが発生することを確認
	assert.Error(t, err)

	// エラーメッセージが適切にラップされていることを確認
	assert.Contains(t, err.Error(), "failed to create task")

	// モックが呼び出されたことを検証
	mockUsecase.AssertExpectations(t)
}
