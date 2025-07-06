package repository

import (
	"database/sql"
	"go-record-app2025/internal/model"
	"log"
)

type PostgresGoalRepository struct {
	DB *sql.DB
}

// インスタンスを生成する関数
// NewPostgresGoalRepositoryは、PostgresGoalRepositoryの新しいインスタンスを
// 作成するための関数です。
func NewPostgresGoalRepository(db *sql.DB) GoalRepository {
	return &PostgresGoalRepository{DB: db}
}

// GetAllは、データベースからすべての目標を取得するメソッドです。
// このメソッドは、目標のステータスごとに分類されたデータを返します。
// 目標のステータスは、"NotStarted", "ActiveGoals", "CompletedGoals"の3つです。
// 返り値は、model.GoalPageData型で、
// 各ステータスごとに目標のスライスを含んでいます。
// エラーが発生した場合は、エラーを返します
func (r *PostgresGoalRepository) GetAllGoals() ([]model.Goal, error) {
	// main.goのsql.Openで取得した接続オブジェクトを使用して、データベースから目標を取得します
	rows, err := r.DB.Query("SELECT id, user_id, title, description, target_date, status FROM goals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []model.Goal
	for rows.Next() {
		var g model.Goal
		if err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Description, &g.TargetDate, &g.Status); err != nil {
			log.Fatal(err)
		}
		goals = append(goals, g)
	}

	return goals, nil

	// ステータスごとに分類
	// var data model.GoalPageData
	// for _, g := range allGoals {
	// 	switch g.Status {
	// 	case "NotStarted":
	// 		data.NotStarted = append(data.NotStarted, g)
	// 	case "ActiveGoals":
	// 		data.ActiveGoals = append(data.ActiveGoals, g)
	// 	case "CompletedGoals":
	// 		data.CompletedGoals = append(data.CompletedGoals, g)
	// 	}
	// }
	// return data, nil
}

func (r *PostgresGoalRepository) SaveGoal(goal model.Goal) error {
	_, err := r.DB.Exec(`
		INSERT INTO goals (user_id, title, description, target_date, status)
		VALUES ($1, $2, $3, $4, $5)
	`, goal.UserID, goal.Title, goal.Description, goal.TargetDate, goal.Status)
	return err
}
