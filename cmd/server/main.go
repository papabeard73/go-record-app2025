package main

import (
	"database/sql"
	"go-record-app2025/internal/handler"
	"go-record-app2025/internal/repository"
	"go-record-app2025/internal/service"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://appuser:apppass@localhost:5432/goalsdb?sslmode=disable")
	if err != nil {
		log.Fatal("DB接続失敗:", err)
	}
	defer db.Close()

	repo := repository.NewPostgresGoalRepository(db)
	svc := service.NewGoalService(repo)
	h := handler.NewGoalHandler(svc)

	http.HandleFunc("/", h.ListGoals)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("🚀 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
