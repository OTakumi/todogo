package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

// init関数でlistコマンドをrootコマンドに登録
func init() {
	rootCmd.AddCommand(listCmd)
}

// listCmd はタスク一覧を表示するコマンドの定義
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `List all tasks in the database.
This command retrieves and displays all tasks with their details including:
- ID
- Title
- Deadline
- Status (Complete/Incomplete)
- Created date`,
	// RunE はlistコマンドのメイン実行関数
	RunE: func(cmd *cobra.Command, args []string) error {
		// データベース操作用のコンテキストを作成
		ctx := context.Background()

		// 共有されているtaskUsecaseインスタンスを使用してすべてのタスクを取得
		// taskUsecaseはmain.goで初期化され、SetupDependencies経由で注入されている
		tasks, err := taskUsecase.FindAll(ctx)
		if err != nil {
			return fmt.Errorf("failed to fetch tasks: %w", err)
		}

		// タスクが存在しない場合の処理
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		// 整形されたテーブル出力のためのtabwriterを作成
		// パラメータ: 出力先, 最小幅, タブ幅, パディング, パディング文字, フラグ
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		// テーブルヘッダーを出力
		fmt.Fprintln(w, "ID\tTitle\tDeadline\tStatus\tCreated")
		fmt.Fprintln(w, "---\t-----\t--------\t------\t-------")

		// 各タスクを反復処理して出力をフォーマット
		for _, task := range tasks {
			// ステータスの表示文字列を決定
			status := "Incomplete"
			if task.IsComplete {
				status = "Complete"
			}

			// 締切日をフォーマット（未設定の場合は"-"を表示）
			deadlineStr := "-"
			if task.Deadline != nil {
				deadlineStr = task.Deadline.Format("2006-01-02")
			}

			// 各タスクの情報を整形された行として出力
			// 各フィールドはタブで区切られ、適切に整列される
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
				task.ID,
				task.Title,
				deadlineStr,
				status,
				task.CreatedAt.Format(time.RFC3339),
			)
		}

		// tabwriterのバッファをフラッシュして、すべての内容を標準出力に書き込む
		if err := w.Flush(); err != nil {
			log.Printf("Warning: failed to flush output: %v", err)
		}

		// タスクの総数を表示
		fmt.Printf("\nTotal: %d task(s)\n", len(tasks))

		return nil
	},
}
