package repository

import "go-record-app2025/internal/model"

type GoalRepository interface {
	GetAllGoals() ([]model.Goal, error)
	SaveGoal(goal model.Goal) error
	GetGoalByID(id int) (model.GoalDetailData, error)
	SaveRecord(record model.StudyRecord) error
	UpdateGoal(goal model.Goal) error
	DeleteGoal(id int) error
	DeleteRecord(id int) error
	GetRecordByID(id int) (model.StudyRecord, error)
	UpdateRecord(record model.StudyRecord) error
}
