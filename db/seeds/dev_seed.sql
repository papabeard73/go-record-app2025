-- 依存関係のあるテーブルから順に削除する必要がある
DELETE FROM study_records;
DELETE FROM goals;
DELETE FROM users;

-- users テーブルの初期データ
INSERT INTO users (id, name, email) VALUES
  (1, 'Eriko', 'eriko@example.com'),
  (2, 'TestUser', 'test@example.com');

-- goals テーブルの初期データ
INSERT INTO goals (id, user_id, title, description, status, target_date) VALUES
  (1, 1, 'Go言語でWebアプリ作成', '学習管理アプリを作成する',  'ActiveGoals', '2025-08-15'),
  (2, 1, 'TOEIC 900点突破', '英語学習を継続して900点を超える', 'NotStarted', '2025-01-01'),
  (3, 2, 'Reactでポートフォリオ作成', '就職活動用にReactアプリを作成する', 'CompletedGoals', '2025-03-31');

-- study_records テーブルの初期データ
INSERT INTO study_records (id, goal_id, content, duration_minutes, recorded_at) VALUES
  (1, 1, 'Go言語の基礎を学習', 60, '2025-01-01 10:00:00'),
  (2, 1, 'Webフレームワークの使い方を学ぶ', 90, '2025-01-02 11:00:00'),
  (3, 2, 'TOEICの問題集を解く', 30, '2025-01-03 14:00:00'),
  (4, 3, 'Reactのコンポーネント設計を学ぶ', 120, '2025-02-01 09:00:00'),
  (5, 3, 'ポートフォリオのデザインを考える', 45, '2025-02-02 15:00:00');
