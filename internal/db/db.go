package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open() *sql.DB {
	env := getenv("APP_ENV", "local")
	var dsn string

	if env == "production" {
		dsn = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require",
			getenv("DB_USER", ""),
			getenv("DB_PASSWORD", ""),
			getenv("DB_HOST", ""),
			getenv("DB_PORT", "5432"),
			getenv("DB_NAME", ""),
		)
	} else {
		dsn = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			getenv("DB_USER", "appuser"),
			getenv("DB_PASSWORD", "apppass"),
			getenv("DB_HOST", "localhost"),
			getenv("DB_PORT", "5432"),
			getenv("DB_NAME", "goalsdb"),
		)
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("DB接続失敗:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("DB疎通失敗:", err)
	}
	return db
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Println("DB切断失敗:", err)
	} else {
		log.Println("DB切断成功")
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
