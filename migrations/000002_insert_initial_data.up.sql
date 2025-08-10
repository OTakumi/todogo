-- 初期データの投入
-- アプリケーションのテスト用サンプルデータを挿入
INSERT INTO tasks (
    id,
    title,
    deadline,
    is_complete,
    created_at,
    updated_at
) VALUES 
(
    'f47ac10b-58cc-4372-a567-0e02b2c3d479',
    'プロジェクトの設計書を作成する',
    '2024-12-31 23:59:59+00:00',
    FALSE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    'e874dcb4-f042-5910-9ab7-0321a60c45ac',
    'データベースのマイグレーションを実装する',
    '2024-12-25 18:00:00+00:00',
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '7d6d370d-a4f1-430b-06c7-d4a363341564',
    'ユニットテストを書く',
    '2024-12-28 12:00:00+00:00',
    FALSE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    'f2f3dbc2-6985-dc25-ecfb-4838a4643e9e',
    'APIドキュメントを更新する',
    NULL,
    FALSE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '99350b0f-65ef-470f-ed59-62808d763b78',
    'コードレビューを完了する',
    '2024-12-24 17:00:00+00:00',
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);
