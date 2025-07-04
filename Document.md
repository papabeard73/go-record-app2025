# 目標・学習管理アプリ
## 要件定義
### 目的の整理
  - 用途：Discordを使っている同好会向け
  - 主機能：Discord内で宣言された目標に対して、個々人が 学習内容・学習時間を記録でき、累積時間を確認したり、自分の取り組みを見える化できる。
  - 要件
    - Go言語
    - 無料 or できるだけ安価で運用したい
    - 小規模から始めて拡張可能にしたい
### 利用者
- 初期想定：Discord同好会メンバーが「個人利用（自分の記録のみ）」を行う
- 将来的に：ログイン or Discord連携によってマルチユーザー化も可能
### 機能
  | 機能名               | 内容                                                                |
  | -------------------- | ------------------------------------------------------------------- |
  | 目標登録             | 自分が取り組んでいる目標を手動で登録（Discord内の宣言を自分で転記） |
  | 学習記録の追加       | 「何を」「何時間やったか」を記録として追加（例：Go言語 1.5時間）    |
  | 累計確認             | 目標ごとの合計時間を算出し表示（「累計を見る」ボタン）              |
  | タイマー機能（任意） | 手動で「記録」を入力せずに、タイマーとして自動記録（後回しでもOK）  |
  | 記録一覧表示         | 各記録を一覧で表示（日付・内容・時間）                              |
### 画面設計（イメージ）
  #### 0. ヘッダー
  - GoalTrack、アイコン
  - Home、Discord、Discordアイコン
  #### 1. 目標一覧画面
  - 自分が登録した目標を一覧表示
    - ステータス：Not Started, Active Goals, Completed Goalsに分けた一覧表示
  - 各目標に対して：
    - 目標タイトル（テキスト）、達成目標日（日付）、累計学習時間（合計値）の表示
    - [学習記録を追加]、[詳細を見る] ボタン
  - 目標新規追加ボタン、累計を見るボタン

  #### 2. 目標追加フォーム
  - 目標タイトル（テキスト）、目標達成日（日付）、目標ステータス（プルダウン）、目標の説明（テキスト）を入力
  - [登録] 、[目標一覧へ戻る] ボタン

  #### 3. 目標詳細画面
  - 目標情報の表示
    - 目標タイトル（テキスト）、目標達成日（日付）、目標の説明（テキスト）
  - [目標の編集]ボタン
  - 紐づく学習記録の一覧表示
    - 各学習記録の [編集] [削除] ボタン
  - 累計学習時間（合計値）の表示
  - 「学習記録を追加」ボタン

  #### 4. 目標編集フォーム
  - 既存の目標情報を編集
    - 目標タイトル（テキスト）、目標達成日（日付）、目標ステータス（プルダウン）、目標の説明（テキスト）を編集
  - [保存] [目標を削除] [目標詳細ページへ戻る] ボタン

  #### 5. 学習記録追加・編集フォーム
  - 日付、内容（テキスト）、学習時間（分単位）を入力
  - [保存] [目標詳細ページへ戻る] ボタン

  #### 6. 累計表示（※目標ごと or 全体）：後回し
  - 集計期間選択（今週、今月、任意期間）
  - 表示方法：棒グラフ、円グラフ
  - フィルター：目標別、ステータス別
  - 合計学習時間（全体）
  #### 補足：安全な削除の方法例
  - 削除前に確認ダイアログを出す
      - 👉「この目標を削除すると関連する学習記録も削除されます。本当によろしいですか？」
  - 「削除」ではなく「アーカイブ」にする方法もあり
    - 👉 間違って消しても復元できる設計に（後で検討）
### 画面遷移
    ```
    [目標一覧]　→　[目標追加]
        ↓ クリック
    [目標詳細・学習記録一覧] → [目標編集]
        ↓
    [学習記録追加・編集]
    ```
### データ構造（Entity案）
- User（将来用。初期はなし）
  - id, name, email, created_at
- Goal（目標）
  - id, title, description, user_id, status（未着手:Not started, 進行中:Active goals, 達成済:Completed goals）, target_date, created_at, updated_at
- StudyRecord（学習記録）
  - id, goal_id, content, duration_minutes(int:分単位), recorded_at
### 可能な拡張機能（将来的に）
  - Discord Bot連携して、メッセージを自動で取り込む（初期は手動でOK）
  - 学習記録のCSV出力
  - カレンダー表示や統計グラフ（学習ペースの可視化）

## 構成
1. 言語
   1. Go言語
      1. DBはPostgreSQL
2. インフラ・運用コストの抑え方
   1. Render
      - [https://render.com](https://render.com/)
      - GoアプリをそのままGitHubと連携してデプロイ可能
      - 無料枠あり（Webサービス：月750時間）
      - 自動スリープあり（アクセスがないときは停止）
3. フロントエンド（UI）
  - シンプルなHTML+JS（Tailwind）
    - デザイン：Stitch（https://stitch.withgoogle.com/）
  - 将来的に：
    - React or Vue + API呼び出し（SPAに）
    - Discord Bot連携（チャットで目標管理など）

## 初期構成イメージ
```
/go-record-app2025
├── cmd/
│   └── server/         ← main.go（エントリーポイント）
│       └── main.go
├── internal/
│   ├── handler/        ← HTTPハンドラ（画面遷移・HTMLレンダリング）
│   ├── service/        ← 業務ロジック（バリデーション・累計計算など）
│   ├── repository/     ← DB操作（SQL or ORM呼び出し）
│   └── model/          ← エンティティ定義（Goal, StudyRecord）
├── templates/          ← HTMLテンプレート（index.htmlなど）
│   ├── layout.html     ← 共通レイアウト
│   ├── goal_list.html
│   ├── goal_detail.html
│   └── ...
├── static/             ← JS/CSSなどの静的ファイル
│   ├── js/
│   │   └── script.js
│   └── css/
│       └── style.css
├── go.mod
└── README.md

```

## 開発の進め方
推奨手順（初期開発）
1. 最小構成でフロントとバックの土台を作る
main.goでHTTPルーティングを設定し、templates/index.htmlが表示されることを確認。
TailwindとJSの読み込み確認（ビルド済みのCDNでも最初はOK）。

2. データベース設計・マイグレーション作成
PostgreSQLのスキーマ作成（gooseやgolang-migrateなどのマイグレーションツールも検討）
GoalとStudyRecordのテーブルを作成。

3. Go側：データ取得・表示の最小ループ
例：トップ画面で「登録された目標一覧」を表示できるようにする（テンプレートに渡す）
Goal一覧取得 → HTMLでループ表示

4. 目標登録フォーム → 登録処理を実装
フォームからPOSTで送信 → DBに登録 → リダイレクト

5. 学習記録の登録・紐づけ表示を実装
学習記録追加 → DB登録 → 目標詳細ページで表示
