package infrastructure

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTaskRepository_FindAll(t *testing.T) {
	// TODO: FindAll test cases
	// - [ ] 空のタスクリストを返すケース
	// - [ ] 複数のタスクを返すケース
	// - [ ] DBエラーが発生するケース

	t.Run("空のタスクリストを返す", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		// tasksテーブルに対するSELECTクエリの期待値を設定する
		// このクエリが実行された際、指定したカラムの行をモックが返す
		mock.ExpectQuery("SELECT id, title, deadline, is_complete, created_at, updated_at FROM tasks").
			WillReturnRows(sqlmock.NewRows([]string{"id", "title", "deadline", "is_complete", "created_at", "updated_at"}))

		// Act
		tasks, err := repo.FindAll(ctx)

		// Assert
		// エラーが発生しないこと
		assert.NoError(t, err)

		// tasksが空であること
		assert.Empty(t, tasks)

		// 設定したモックの期待値が全て満たされること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("複数のタスクを返す", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		now := time.Now()
		rows := sqlmock.NewRows([]string{"id", "title", "deadline", "is_complete", "created_at", "updated_at"}).
			AddRow("1", "Task 1", now, false, now, now).
			AddRow("2", "Task 2", now.Add(24*time.Hour), true, now, now)

		mock.ExpectQuery("SELECT id, title, deadline, is_complete, created_at, updated_at FROM tasks").
			WillReturnRows(rows)

		// Act
		tasks, err := repo.FindAll(ctx)

		// Assert
		// エラーが発生しないこと
		assert.NoError(t, err)

		// tasksが2件のタスクを含んでいること
		assert.Len(t, tasks, 2)

		// タスクの内容が、登録した内容と一致すること
		assert.Equal(t, "1", tasks[0].ID)
		assert.Equal(t, "Task 1", tasks[0].Title)
		assert.False(t, tasks[0].IsComplete)
		assert.Equal(t, "2", tasks[1].ID)
		assert.Equal(t, "Task 2", tasks[1].Title)
		assert.True(t, tasks[1].IsComplete)

		// 設定したモックの期待値が全て満たされること
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DBエラーが発生する", func(t *testing.T) {
		// Arrange
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		repo := NewTaskRepository(db)
		ctx := context.Background()

		mock.ExpectQuery("SELECT id, title, deadline, is_complete, created_at, updated_at FROM tasks").
			WillReturnError(sql.ErrConnDone)

		// Act
		tasks, err := repo.FindAll(ctx)

		// Assert
		// エラーが発生すること
		assert.Error(t, err)
		assert.Nil(t, tasks)

		// モックがエラーを返すこと
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

