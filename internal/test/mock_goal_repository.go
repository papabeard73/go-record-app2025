package test

import (
	"go-record-app2025/internal/model"
)

type MockGoalRepo struct {
	SaveRecordCalled bool
	SavedRecord      model.StudyRecord
}

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

func (m *MockGoalRepo) MockGoalData() error {
	return nil
}

func (m *MockGoalRepo) GetGoalByID(id int) (model.GoalDetailData, error) {
	// テスト用のダミー返却
	return model.GoalDetailData{
		Goal: model.Goal{
			ID:          id,
			UserID:      1,
			Title:       "Mock Goal",
			Description: "This is a mock goal for testing.",
			TargetDate:  "2025-12-31",
			Status:      "ActiveGoals",
		},
		StudyRecords: []model.StudyRecord{
			{ID: 1, GoalID: id, Content: "Study content", DurationMinutes: 30, RecordedAt: "2025-01-01T10:00:00Z"},
		},
	}, nil
}

// MockGoalRepoは、GoalRepositoryインターフェースのモック実装です。
func (m *MockGoalRepo) SaveRecord(record model.StudyRecord) error {
	m.SaveRecordCalled = true
	m.SavedRecord = record
	return nil
}

func MockStudyRecord() model.StudyRecord {
	return model.StudyRecord{
		ID:              1,
		GoalID:          1,
		Content:         "Mock Study Record",
		DurationMinutes: 30,
		RecordedAt:      "2025-01-01T10:00:00Z",
	}
}

func (m *MockGoalRepo) DeleteGoal(id int) error {
	// テスト用だから何もしなくてOK！
	return nil
}
func (m *MockGoalRepo) UpdateGoal(goal model.Goal) error {
	// テスト用なので何もしない or 必要ならモックの挙動を書く
	return nil
}
