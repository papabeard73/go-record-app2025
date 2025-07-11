package service

import (
	"go-record-app2025/internal/model"
	"go-record-app2025/internal/repository"
)

// GoalServiceは、目標に関連するビジネスロジックを提供するサービスです。
// このサービスは、目標のデータを取得するためのメソッドを
// 提供します。
// 目標のデータは、リポジトリパターンを使用して
// データベースから取得されます。
// GoalServiceは、目標の一覧を取得するためのGetGoalsメソッドを持っています。
// このメソッドは、リポジトリから目標のデータを取得し、
// model.GoalPageData型で返します。
// 目標のデータは、ステータスごとに分類されており、
// "NotStarted", "ActiveGoals", "CompletedGoals"の3つのステータスがあります。
// それぞれのステータスに対応する目標のスライスが
// model.GoalPageData型のフィールドとして定義されています。
type GoalService struct {
	Repo repository.GoalRepository
}

// NewGoalServiceは、GoalServiceの新しいインスタンスを作成するための関数です。
func NewGoalService(repo repository.GoalRepository) *GoalService {
	return &GoalService{Repo: repo}
}

// GetGoalsは、目標の一覧を取得するメソッドです。
func (s *GoalService) GetGoals() (model.GoalPageData, error) {
	goals, err := s.Repo.GetAllGoals()
	if err != nil {
		return model.GoalPageData{}, err
	}

	var pageData model.GoalPageData
	for _, g := range goals {
		switch g.Status {
		case "NotStarted":
			pageData.NotStarted = append(pageData.NotStarted, g)
		case "ActiveGoals":
			pageData.ActiveGoals = append(pageData.ActiveGoals, g)
		case "CompletedGoals":
			pageData.CompletedGoals = append(pageData.CompletedGoals, g)
		}
	}
	return pageData, nil

	// return s.Repo.GetAll()
}

// CreateGoalは、新しい目標を追加するメソッドです。
// このメソッドは、目標のデータを受け取り、リポジトリを通じて
// データベースに保存します。
func (s *GoalService) CreateGoal(goal model.Goal) error {
	return s.Repo.SaveGoal(goal)
}
