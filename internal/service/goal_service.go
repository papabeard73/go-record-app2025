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

// DetailGoalsは、特定の目標の詳細を取得するメソッドです。
// このメソッドは、目標のIDを受け取り、リポジトリを通じて
// データベースから目標の詳細を取得します。
func (s *GoalService) DetailGoals(id int) (model.GoalDetailData, error) {
	goal, err := s.Repo.GetGoalByID(id)
	if err != nil {
		return model.GoalDetailData{}, err
	}
	return goal, nil
}

// AddNewRecordは、目標に新しい学習記録を追加するメソッドです。
// このメソッドは、学習記録のデータを受け取り、
// リポジトリを通じてデータベースに保存します。
func (s *GoalService) AddNewRecord(record model.StudyRecord) error {
	return s.Repo.SaveRecord(record)
}

func (s *GoalService) UpdateGoal(goal model.Goal) error {
	return s.Repo.UpdateGoal(goal)
}

// DeleteGoalは、目標を削除するメソッドです。
// このメソッドは、目標のIDを受け取り、リポジトリを通じて
// データベースから目標を削除します。
func (s *GoalService) DeleteGoal(id int) error {
	return s.Repo.DeleteGoal(id)
}

// DeleteRecordは、学習記録を削除するメソッドです。
// このメソッドは、学習記録のIDを受け取り、リポジトリを通じて
// データベースから学習記録を削除します。
func (s *GoalService) DeleteRecord(id int) error {
	return s.Repo.DeleteRecord(id)
}

func (s *GoalService) GetRecordByID(id int) (model.StudyRecord, error) {
	return s.Repo.GetRecordByID(id)
}

func (s *GoalService) UpdateRecord(record model.StudyRecord) error {
	return s.Repo.UpdateRecord(record)
}
