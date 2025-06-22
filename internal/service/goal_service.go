package service

import (
	"go-record-app2025/internal/model"
	"go-record-app2025/internal/repository"
)

type GoalService struct {
	Repo repository.GoalRepository
}

func NewGoalService(repo repository.GoalRepository) *GoalService {
	return &GoalService{Repo: repo}
}

func (s *GoalService) GetGoals() (model.GoalPageData, error) {
	return s.Repo.GetAll()
}
