-- tasksテーブルの作成
-- タスク管理アプリケーションのメインテーブル
CREATE TABLE IF NOT EXISTS tasks (
    -- 主キー: UUID形式のタスクID
    id VARCHAR(36) PRIMARY KEY,
    
    -- タスクのタイトル（必須）
    title VARCHAR(255) NOT NULL,
    
    -- タスクの締切日時（NULL許可）
    deadline TIMESTAMP WITH TIME ZONE,
    
    -- タスクの完了状態（デフォルトは未完了）
    is_complete BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- レコードの作成日時
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- レコードの更新日時
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 更新日時を自動的に更新するための関数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 更新日時自動更新のトリガー
CREATE TRIGGER update_tasks_updated_at BEFORE UPDATE
    ON tasks FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- インデックスの作成
-- 締切日でのソートや検索を高速化
CREATE INDEX idx_tasks_deadline ON tasks(deadline);

-- 完了状態での絞り込みを高速化
CREATE INDEX idx_tasks_is_complete ON tasks(is_complete);

-- 作成日時でのソートを高速化
CREATE INDEX idx_tasks_created_at ON tasks(created_at);