package repository

import (
	"database/sql"
	"go-record-app2025/internal/model"
	"log"
)

type PostgresGoalRepository struct {
	DB *sql.DB
}

func NewPostgresGoalRepository(db *sql.DB) GoalRepository {
	return &PostgresGoalRepository{DB: db}
}

func (r *PostgresGoalRepository) GetAll() (model.GoalPageData, error) {
	rows, err := r.DB.Query("SELECT id, user_id, title, description, target_date, status FROM goals")
	if err != nil {
		return model.GoalPageData{}, err
	}
	defer rows.Close()

	// var goals []model.Goal
	// for rows.Next() {
	// 	var g model.Goal
	// 	if err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Description, &g.TargetDate, &g.Status); err != nil {
	// 		return nil, err
	// 	}
	// 	goals = append(goals, g)
	// }
	var allGoals []model.Goal
	for rows.Next() {
		var g model.Goal
		if err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Description, &g.TargetDate, &g.Status); err != nil {
			log.Fatal(err)
		}
		allGoals = append(allGoals, g)
	}

	// ステータスごとに分類
	var data model.GoalPageData
	for _, g := range allGoals {
		switch g.Status {
		case "NotStarted":
			data.NotStarted = append(data.NotStarted, g)
		case "ActiveGoals":
			data.ActiveGoals = append(data.ActiveGoals, g)
		case "CompletedGoals":
			data.CompletedGoals = append(data.CompletedGoals, g)
		}
	}
	return data, nil
}
