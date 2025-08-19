package infrastructure

import (
	"OTakumi/todogo/internal/domain/model"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// TestTaskRepository_Create はTaskRepositoryのCreateメソッドのテストケース
// TODO: Create test cases
// - [ ] 正常にタスクを作成できるケース
// - [ ] タイトルが空文字の場合にバリデーションエラーが発生するケース
// - [ ] Deadlineがnilの場合でも正常に作成できるケース
// - [ ] Deadlineが設定されている場合も正常に作成できるケース
// - [ ] DeadlineがCreatedAtよりも過去の場合にバリデーションエラーが発生するケース
// - [ ] DBへのINSERTが失敗するケース
// - [ ] IDの自動生成が正しく動作するケース
// - [ ] CreatedAtとUpdatedAtが正しく設定されるケース
func TestTaskRepository_Create(t *testing.T) {
	t.Run("正常にタスクを作成できる", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		// テスト用のタスクを作成
		// IDは自動生成されることを想定してnewTaskではIDを設定しない
		newTask := &model.Task{
			Title:      "新しいタスク",
			Deadline:   nil,
			IsComplete: false,
		}

		// トランザクションの期待値を設定
		mock.ExpectBegin()
		// INSERTクエリの期待値を設定
		// IDはUUIDで自動生成されることを想定
		mock.ExpectExec("INSERT INTO tasks").
			WithArgs(
				sqlmock.AnyArg(), // ID (UUID)
				"新しいタスク",         // Title
				nil,              // Deadline
				false,            // IsComplete
				sqlmock.AnyArg(), // CreatedAt
				sqlmock.AnyArg(), // UpdatedAt
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Act
		createdTask, err := repo.Create(ctx, newTask)

		// Assert
		// エラーが発生しないこと
		assert.NoError(t, err)
		assert.NotNil(t, createdTask)

		// 作成されたタスクの内容を検証
		assert.NotEmpty(t, createdTask.ID) // IDが自動生成されていること
		assert.Equal(t, "新しいタスク", createdTask.Title)
		assert.Nil(t, createdTask.Deadline)
		assert.False(t, createdTask.IsComplete)
		assert.NotZero(t, createdTask.CreatedAt) // CreatedAtが設定されていること
		assert.NotZero(t, createdTask.UpdatedAt) // UpdatedAtが設定されていること

		// モックの期待値が満たされていること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("タイトルが空文字の場合にバリデーションエラーが発生する", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		// タイトルが空のタスクを作成
		invalidTask := &model.Task{
			Title:      "", // 空文字
			Deadline:   nil,
			IsComplete: false,
		}

		// DBクエリは実行されないことを期待（バリデーションで弾かれるため）

		// Act
		createdTask, err := repo.Create(ctx, invalidTask)

		// Assert
		// バリデーションエラーが発生すること
		assert.Error(t, err)
		assert.Nil(t, createdTask)
		assert.Contains(t, err.Error(), "Title is required")

		// モックの期待値が満たされていること（クエリが実行されていないこと）
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Deadlineが設定されている場合も正常に作成できる", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		// 期限付きのタスクを作成
		deadline := time.Now().Add(7 * 24 * time.Hour) // 1週間後
		taskWithDeadline := &model.Task{
			Title:      "期限付きタスク",
			Deadline:   &deadline,
			IsComplete: false,
		}

		// トランザクションの期待値を設定
		mock.ExpectBegin()
		// INSERTクエリの期待値を設定
		mock.ExpectExec("INSERT INTO tasks").
			WithArgs(
				sqlmock.AnyArg(), // ID (UUID)
				"期限付きタスク",        // Title
				deadline,         // Deadline
				false,            // IsComplete
				sqlmock.AnyArg(), // CreatedAt
				sqlmock.AnyArg(), // UpdatedAt
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Act
		createdTask, err := repo.Create(ctx, taskWithDeadline)

		// Assert
		// エラーが発生しないこと
		assert.NoError(t, err)
		assert.NotNil(t, createdTask)

		// 作成されたタスクの内容を検証
		assert.NotEmpty(t, createdTask.ID)
		assert.Equal(t, "期限付きタスク", createdTask.Title)
		assert.NotNil(t, createdTask.Deadline)
		assert.Equal(t, deadline.Unix(), createdTask.Deadline.Unix()) // 秒単位で比較
		assert.False(t, createdTask.IsComplete)

		// モックの期待値が満たされていること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DBへのINSERTが失敗する", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		task := &model.Task{
			Title:      "エラーテスト用タスク",
			Deadline:   nil,
			IsComplete: false,
		}

		// トランザクションの期待値を設定
		mock.ExpectBegin()
		// INSERTクエリがエラーを返すように設定
		mock.ExpectExec("INSERT INTO tasks").
			WithArgs(
				sqlmock.AnyArg(), // ID
				"エラーテスト用タスク",     // Title
				nil,              // Deadline
				false,            // IsComplete
				sqlmock.AnyArg(), // CreatedAt
				sqlmock.AnyArg(), // UpdatedAt
			).
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		// Act
		createdTask, err := repo.Create(ctx, task)

		// Assert
		// エラーが発生すること
		assert.Error(t, err)
		assert.Nil(t, createdTask)
		assert.Contains(t, err.Error(), "failed to insert task")

		// モックの期待値が満たされていること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("IDの自動生成が正しく動作する", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		// IDが設定されていないタスクを作成
		taskWithoutID := &model.Task{
			ID:         "", // 空のID
			Title:      "ID自動生成テスト",
			Deadline:   nil,
			IsComplete: false,
		}

		// トランザクションの期待値を設定
		mock.ExpectBegin()
		// INSERTクエリの期待値を設定
		// IDはUUID形式で自動生成されることを期待
		mock.ExpectExec("INSERT INTO tasks").
			WithArgs(
				sqlmock.AnyArg(), // 自動生成されたID
				"ID自動生成テスト",      // Title
				nil,              // Deadline
				false,            // IsComplete
				sqlmock.AnyArg(), // CreatedAt
				sqlmock.AnyArg(), // UpdatedAt
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Act
		createdTask, err := repo.Create(ctx, taskWithoutID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, createdTask)

		// IDが自動生成されていること（UUID形式の検証）
		assert.NotEmpty(t, createdTask.ID)
		assert.Len(t, createdTask.ID, 36)       // UUID v4の標準的な長さ
		assert.Contains(t, createdTask.ID, "-") // UUIDにはハイフンが含まれる

		// モックの期待値が満たされていること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("CreatedAtとUpdatedAtが正しく設定される", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		// 現在時刻を記録（タスク作成前）
		beforeCreate := time.Now()

		task := &model.Task{
			Title:      "タイムスタンプテスト",
			Deadline:   nil,
			IsComplete: false,
		}

		// トランザクションの期待値を設定
		mock.ExpectBegin()
		// INSERTクエリの期待値を設定
		mock.ExpectExec("INSERT INTO tasks").
			WithArgs(
				sqlmock.AnyArg(), // ID
				"タイムスタンプテスト",     // Title
				nil,              // Deadline
				false,            // IsComplete
				sqlmock.AnyArg(), // CreatedAt
				sqlmock.AnyArg(), // UpdatedAt
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Act
		createdTask, err := repo.Create(ctx, task)

		// 現在時刻を記録（タスク作成後）
		afterCreate := time.Now()

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, createdTask)

		// CreatedAtとUpdatedAtが設定されていること
		assert.NotZero(t, createdTask.CreatedAt)
		assert.NotZero(t, createdTask.UpdatedAt)

		// タイムスタンプが妥当な範囲内であること
		assert.True(t, createdTask.CreatedAt.After(beforeCreate) || createdTask.CreatedAt.Equal(beforeCreate))
		assert.True(t, createdTask.CreatedAt.Before(afterCreate) || createdTask.CreatedAt.Equal(afterCreate))
		assert.True(t, createdTask.UpdatedAt.After(beforeCreate) || createdTask.UpdatedAt.Equal(beforeCreate))
		assert.True(t, createdTask.UpdatedAt.Before(afterCreate) || createdTask.UpdatedAt.Equal(afterCreate))

		// CreatedAtとUpdatedAtが同じ値であること（新規作成時）
		assert.Equal(t, createdTask.CreatedAt.Unix(), createdTask.UpdatedAt.Unix())

		// モックの期待値が満たされていること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("トランザクション開始に失敗する", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		task := &model.Task{
			Title:      "トランザクションエラーテスト",
			Deadline:   nil,
			IsComplete: false,
		}

		// トランザクション開始でエラーを返すように設定
		mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

		// Act
		createdTask, err := repo.Create(ctx, task)

		// Assert
		// エラーが発生すること
		assert.Error(t, err)
		assert.Nil(t, createdTask)
		assert.Contains(t, err.Error(), "failed to begin transaction")

		// モックの期待値が満たされていること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("コミットに失敗する", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		task := &model.Task{
			Title:      "コミットエラーテスト",
			Deadline:   nil,
			IsComplete: false,
		}

		// トランザクションの期待値を設定
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO tasks").
			WithArgs(
				sqlmock.AnyArg(), // ID
				"コミットエラーテスト",     // Title
				nil,              // Deadline
				false,            // IsComplete
				sqlmock.AnyArg(), // CreatedAt
				sqlmock.AnyArg(), // UpdatedAt
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		// コミットでエラーを返すように設定
		mock.ExpectCommit().WillReturnError(sql.ErrTxDone)

		// Act
		createdTask, err := repo.Create(ctx, task)

		// Assert
		// エラーが発生すること
		assert.Error(t, err)
		assert.Nil(t, createdTask)
		assert.Contains(t, err.Error(), "failed to commit transaction")

		// モックの期待値が満たされていること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DeadlineがCreatedAtよりも過去の場合にバリデーションエラーが発生する", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer func() { _ = db.Close() }()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		// 過去の期限を設定したタスクを作成
		pastDeadline := time.Now().Add(-24 * time.Hour) // 1日前
		taskWithPastDeadline := &model.Task{
			Title:      "過去の期限タスク",
			Deadline:   &pastDeadline,
			IsComplete: false,
		}

		// DBクエリは実行されないことを期待（バリデーションで弾かれるため）

		// Act
		createdTask, err := repo.Create(ctx, taskWithPastDeadline)

		// Assert
		// バリデーションエラーが発生すること
		assert.Error(t, err)
		assert.Nil(t, createdTask)
		assert.Contains(t, err.Error(), "Deadline must be in the future")

		// モックの期待値が満たされていること（クエリが実行されていないこと）
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
