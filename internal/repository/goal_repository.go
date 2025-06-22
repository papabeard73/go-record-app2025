package repository

import "go-record-app2025/internal/model"

type GoalRepository interface {
	GetAll() (model.GoalPageData, error)
}
