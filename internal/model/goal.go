package model

type Goal struct {
	ID          int
	UserID      int
	Title       string
	Description string
	TargetDate  string
	Status      string // "NotStarted", "ActiveGoals", "CompletedGoals"
}

type GoalPageData struct {
	NotStarted     []Goal
	ActiveGoals    []Goal
	CompletedGoals []Goal
}
