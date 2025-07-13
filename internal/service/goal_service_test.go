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
