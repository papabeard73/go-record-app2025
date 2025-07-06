package test

import (
	"go-record-app2025/internal/model"
)

type MockGoalRepo struct{}

func (m *MockGoalRepo) GetAllGoals() ([]model.Goal, error) {
	return []model.Goal{
		{ID: 1, UserID: 1, Title: "Test Goal 1", Description: "Description 1", TargetDate: "2025-12-31", Status: "NotStarted"},
		{ID: 2, UserID: 1, Title: "Test Goal 2", Description: "Description 2", TargetDate: "2025-12-31", Status: "ActiveGoals"},
		{ID: 3, UserID: 1, Title: "Test Goal 3", Description: "Description 3", TargetDate: "2025-12-31", Status: "CompletedGoals"},
	}, nil
}

func (m *MockGoalRepo) SaveGoal(goal model.Goal) error {
	// テスト用なので何もしない or 必要ならモックの挙動を書く
	return nil
}
