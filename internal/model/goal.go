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

type StudyRecord struct {
	ID              int
	GoalID          int
	Content         string
	DurationMinutes int // 分単位
	RecordedAt      string
}
type GoalDetailData struct {
	Goal Goal
	// 他の必要なフィールドを追加できます
	StudyRecords []StudyRecord
}

type AddRecordPageData struct {
	GoalID    int
	GoalTitle string
}
