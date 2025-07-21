package service_test

import (
	"testing"

	"go-record-app2025/internal/service"
	"go-record-app2025/internal/test"

	"github.com/stretchr/testify/assert"
)

func TestGetGoals(t *testing.T) {
	mockRepo := &test.MockGoalRepo{}
	svc := service.NewGoalService(mockRepo)

	goals, err := svc.GetGoals()

	assert.NoError(t, err)
	assert.Len(t, goals.NotStarted, 1)
	assert.Equal(t, "モック目標1", goals.NotStarted[0].Title)
}

func TestDetailGoals(t *testing.T) {
	mockRepo := &test.MockGoalRepo{}
	svc := service.NewGoalService(mockRepo)

	goal, err := svc.DetailGoals(1)

	assert.NoError(t, err)
	assert.Equal(t, "Mock Goal", goal.Goal.Title)
	assert.Equal(t, "This is a mock goal for testing.", goal.Goal.Description)
	assert.Len(t, goal.StudyRecords, 1)
	assert.Equal(t, "Study content", goal.StudyRecords[0].Content)
	assert.Equal(t, 30, goal.StudyRecords[0].DurationMinutes)
}
func TestAddNewRecord(t *testing.T) {
	mockRepo := &test.MockGoalRepo{}
	svc := service.NewGoalService(mockRepo)

	record := test.MockStudyRecord() // モックのStudyRecordを取得する関数がtestパッケージにある前提

	err := svc.AddNewRecord(record)

	assert.NoError(t, err)
	assert.True(t, mockRepo.SaveRecordCalled, "SaveRecord should be called")
	assert.Equal(t, record, mockRepo.SavedRecord, "Saved record should match input")
}
