package model_test

import (
	"OTakumi/todogo/internal/domain/model"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	t.Run("NewTask関数がTaskを問題なく生成すること", func(t *testing.T) {
		// Arrange
		id := "a95fbaab-5356-94e9-011c-97b0e37af5aa"
		title := "Testing Go"

		// Act
		task := model.NewTask(id, title)

		// Assert
		// オブジェクトがnilでないこと
		if task == nil {
			t.Fatal("NewTask() returned nil")
		}

		// Titleが引数の通りに設定されていること
		if task.Title != title {
			t.Errorf("expected Title to be '%s', but got '%s'", title, task.Title)
		}

		// 作成直後なので、IsCompleteがfalseになっていること
		if task.IsComplete != false {
			t.Errorf("expected Completed to be false, but got %v", task.IsComplete)
		}

		// CreatedAtに作成日時が設定されていること
		if task.CreatedAt.IsZero() || time.Since(task.CreatedAt) > time.Second {
			t.Errorf("expected CreatedAt to be set to the current time, but got %v", task.CreatedAt)
		}

		// UpdatedAtに作成日時が設定されていること
		if task.UpdatedAt.IsZero() || time.Since(task.UpdatedAt) > time.Second {
			t.Errorf("expected UpdatedAt to be set to the current time, but got %v", task.UpdatedAt)
		}
	})
}

func TestTask_Validate(t *testing.T) {
	t.Run("タスクのタイトルが空の場合、エラーが返されること", func(t *testing.T) {
		// Arrange
		task := model.Task{
			Title: "",
		}

		// Act
		err := task.Validate()

		// Assert
		// Validateがエラーを返す
		// nilの場合、テストが失敗する
		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})

	t.Run("タスクのタイトルが空でない場合、nilが返されること", func(t *testing.T) {
		// Arrange
		task := model.Task{
			Title: "Test Task",
		}

		// Act
		err := task.Validate()
		// Assert
		// Validateがnilを返す
		if err != nil {
			// エラーの中身を確認する
			t.Errorf("did not expect an error, but got: %v", err)
		}
	})
}
