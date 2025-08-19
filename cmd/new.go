package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// titleフラグの値を格納する変数
var taskTitle string

func init() {
	// newコマンドをrootコマンドに追加
	rootCmd.AddCommand(newCmd)

	// newコマンドにtitleフラグを追加
	// このフラグは必須で、タスクのタイトルを指定するために使用される
	newCmd.Flags().StringVarP(&taskTitle, "title", "t", "", "Task title (required)")
	newCmd.MarkFlagRequired("title")
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new task",
	Long: `Create a new task with a title.
	
This command creates a new task in the database with the specified title.
The task will be created with default values for other fields.`,
	// RunEを使用してエラーハンドリングを可能にする
	RunE: func(cmd *cobra.Command, args []string) error {
		// タイトルが空の場合のバリデーション
		if taskTitle == "" {
			return errors.New("title cannot be empty")
		}

		// コンテキストの作成（タイムアウトやキャンセレーション用）
		ctx := context.Background()

		// Usecaseレイヤーを使用してタスクを作成
		// taskUsecaseはroot.goで定義され、SetupDependencies関数で初期化される
		createdTask, err := taskUsecase.CreateTask(ctx, taskTitle)
		if err != nil {
			// エラーをラップして上位層に返す
			return fmt.Errorf("failed to create task: %w", err)
		}

		// 成功メッセージを出力
		fmt.Fprintf(cmd.OutOrStdout(), "Task created successfully!\n")
		fmt.Fprintf(cmd.OutOrStdout(), "ID: %s\n", createdTask.ID)
		fmt.Fprintf(cmd.OutOrStdout(), "Title: %s\n", createdTask.Title)

		return nil
	},
}
