package repository

import "go-record-app2025/internal/model"

type GoalRepository interface {
	GetAllGoals() ([]model.Goal, error)
	SaveGoal(goal model.Goal) error
}
