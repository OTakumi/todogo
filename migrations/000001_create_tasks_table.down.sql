-- マイグレーションのロールバック用SQL
-- テーブル作成の逆操作を定義

-- トリガーの削除
DROP TRIGGER IF EXISTS update_tasks_updated_at ON tasks;

-- 関数の削除
DROP FUNCTION IF EXISTS update_updated_at_column();

-- インデックスの削除
DROP INDEX IF EXISTS idx_tasks_created_at;
DROP INDEX IF EXISTS idx_tasks_is_complete;
DROP INDEX IF EXISTS idx_tasks_deadline;

-- tasksテーブルの削除
DROP TABLE IF EXISTS tasks;