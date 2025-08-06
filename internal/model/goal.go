package model

import "time"

type Goal struct {
	ID            int
	UserID        int
	Title         string
	Description   string
	TargetDate    time.Time
	TargetDateStr string // フォーマット済みのターゲット日付文字列（表示用）
	Status        string // "NotStarted", "ActiveGoals", "CompletedGoals"
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
	RecordedAt      time.Time
	RecordedAtStr   string // フォーマット済みの記録日時文字列（表示用）
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
