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
