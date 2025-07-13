package repository

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetGoalByID_Found(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := &PostgresGoalRepository{DB: db}

	goalID := 1
	goalRows := sqlmock.NewRows([]string{"id", "user_id", "title", "description", "target_date", "status"}).
		AddRow(goalID, 2, "Test Title", "Test Desc", "2024-06-01", "ActiveGoals")
	mock.ExpectQuery("SELECT id, user_id, title, description, target_date, status FROM goals WHERE id = \\$1").
		WithArgs(goalID).WillReturnRows(goalRows)

	recordRows := sqlmock.NewRows([]string{"id", "goal_id", "content", "duration_minutes", "recorded_at"}).
		AddRow(10, goalID, "Study Go", 60, "2024-06-02")
	mock.ExpectQuery("SELECT id, goal_id, content, duration_minutes, recorded_at FROM study_records WHERE goal_id = \\$1").
		WithArgs(goalID).WillReturnRows(recordRows)

	result, err := repo.GetGoalByID(goalID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Goal.ID != goalID {
		t.Errorf("expected goal ID %d, got %d", goalID, result.Goal.ID)
	}
	if len(result.StudyRecords) != 1 {
		t.Errorf("expected 1 study record, got %d", len(result.StudyRecords))
	}
	if result.StudyRecords[0].Content != "Study Go" {
		t.Errorf("expected study record content 'Study Go', got '%s'", result.StudyRecords[0].Content)
	}
}

func TestGetGoalByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := &PostgresGoalRepository{DB: db}

	goalID := 99
	mock.ExpectQuery("SELECT id, user_id, title, description, target_date, status FROM goals WHERE id = \\$1").
		WithArgs(goalID).WillReturnError(sql.ErrNoRows)

	result, err := repo.GetGoalByID(goalID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Goal.ID != 0 {
		t.Errorf("expected empty goal, got ID %d", result.Goal.ID)
	}
}

func TestGetGoalByID_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := &PostgresGoalRepository{DB: db}

	goalID := 1
	mock.ExpectQuery("SELECT id, user_id, title, description, target_date, status FROM goals WHERE id = \\$1").
		WithArgs(goalID).WillReturnError(errors.New("db error"))

	_, err = repo.GetGoalByID(goalID)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetGoalByID_StudyRecordError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := &PostgresGoalRepository{DB: db}

	goalID := 1
	goalRows := sqlmock.NewRows([]string{"id", "user_id", "title", "description", "target_date", "status"}).
		AddRow(goalID, 2, "Test Title", "Test Desc", "2024-06-01", "ActiveGoals")
	mock.ExpectQuery("SELECT id, user_id, title, description, target_date, status FROM goals WHERE id = \\$1").
		WithArgs(goalID).WillReturnRows(goalRows)

	mock.ExpectQuery("SELECT id, goal_id, content, duration_minutes, recorded_at FROM study_records WHERE goal_id = \\$1").
		WithArgs(goalID).WillReturnError(errors.New("study record error"))

	_, err = repo.GetGoalByID(goalID)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
