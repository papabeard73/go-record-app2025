package test

import (
	"go-record-app2025/internal/model"
)

type MockGoalRepo struct{}

func (m *MockGoalRepo) GetAll() (model.GoalPageData, error) {
	return model.GoalPageData{
		NotStarted: []model.Goal{
			{ID: 1, Title: "モック目標1", Status: "NotStarted"},
		},
		ActiveGoals: []model.Goal{
			{ID: 2, Title: "モック目標2", Status: "ActiveGoals"},
		},
		CompletedGoals: []model.Goal{},
	}, nil
}
