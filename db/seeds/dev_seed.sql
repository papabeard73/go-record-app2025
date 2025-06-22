-- 依存関係のあるテーブルから順に削除する必要がある
DELETE FROM goals;
DELETE FROM users;

-- users テーブルの初期データ
INSERT INTO users (id, name, email) VALUES
  (1, 'Eriko', 'eriko@example.com'),
  (2, 'TestUser', 'test@example.com');

-- goals テーブルの初期データ
INSERT INTO goals (user_id, title, description, target_date, status) VALUES
  (1, 'Go言語でWebアプリ作成', '学習管理アプリを作成する', '2025-08-15', 'ActiveGoals'),
  (1, 'TOEIC 900点突破', '英語学習を継続して900点を超える', '2025-01-01', 'NotStarted'),
  (2, 'Reactでポートフォリオ作成', '就職活動用にReactアプリを作成する', '2025-03-31', 'CompletedGoals');

