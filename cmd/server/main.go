package main

import (
	"go-record-app2025/internal/db"
	"go-record-app2025/internal/handler"
	"go-record-app2025/internal/repository"
	"go-record-app2025/internal/service"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQLデータベースに接続
	// データベースURLは環境変数や設定ファイルから取得する
	// ここでは直接指定していますが、実際のアプリケーションでは環境変数や設定ファイルから取得することを推奨します。
	// 注意: 本番環境ではSSL接続を使用することを推奨します。
	// ここではSSLモードを無効にしていますが、実際の運用では適切なSSL設定を行ってください。
	// 例: "postgresql://appuser:apppass@localhost:5432/goalsdb?sslmode=require"
	// ここでは、PostgreSQLのドライバを使用しています。
	// ドライバのインポートは、_ "github.com/lib/pq"
	// のようにアンダースコアで行うことで、パッケージの初期化を行います。
	// このパッケージは、PostgreSQLのドライバを提供します。
	// db, err := sql.Open("postgres", "postgresql://appuser:apppass@localhost:5432/goalsdb?sslmode=disable")
	database := db.Open()
	// if err != nil {
	// 	log.Fatal("DB接続失敗:", err)
	// }
	defer db.Close(database)

	// データベースの接続を確認
	repo := repository.NewPostgresGoalRepository(database)
	// リポジトリを使用してサービスを初期化
	svc := service.NewGoalService(repo)
	// ハンドラーを初期化
	// ハンドラーは、HTTPリクエストを処理するためのもの
	h := handler.NewGoalHandler(svc)

	// HTTPサーバーを設定
	http.HandleFunc("/", h.ListGoals)
	http.HandleFunc("/goals/new", h.AddNewGoals)
	http.HandleFunc("/goals/detail", h.DetailGoals)
	http.HandleFunc("/records/new", h.AddNewRecord)
	//
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("🚀 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
