package repository

import (
	"database/sql"
	"go-record-app2025/internal/model"
	"log"
)

type PostgresGoalRepository struct {
	DB *sql.DB
}

// インスタンスを生成する関数
// NewPostgresGoalRepositoryは、PostgresGoalRepositoryの新しいインスタンスを
// 作成するための関数です。
func NewPostgresGoalRepository(db *sql.DB) GoalRepository {
	return &PostgresGoalRepository{DB: db}
}

// GetAllは、データベースからすべての目標を取得するメソッドです。
// このメソッドは、目標のステータスごとに分類されたデータを返します。
// 目標のステータスは、"NotStarted", "ActiveGoals", "CompletedGoals"の3つです。
// 返り値は、model.GoalPageData型で、
// 各ステータスごとに目標のスライスを含んでいます。
// エラーが発生した場合は、エラーを返します
func (r *PostgresGoalRepository) GetAllGoals() ([]model.Goal, error) {
	// main.goのsql.Openで取得した接続オブジェクトを使用して、データベースから目標を取得します
	rows, err := r.DB.Query("SELECT id, user_id, title, description, target_date, status FROM goals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []model.Goal
	for rows.Next() {
		var g model.Goal
		if err := rows.Scan(&g.ID, &g.UserID, &g.Title, &g.Description, &g.TargetDate, &g.Status); err != nil {
			log.Fatal(err)
		}
		goals = append(goals, g)
	}

	return goals, nil

	// ステータスごとに分類
	// var data model.GoalPageData
	// for _, g := range allGoals {
	// 	switch g.Status {
	// 	case "NotStarted":
	// 		data.NotStarted = append(data.NotStarted, g)
	// 	case "ActiveGoals":
	// 		data.ActiveGoals = append(data.ActiveGoals, g)
	// 	case "CompletedGoals":
	// 		data.CompletedGoals = append(data.CompletedGoals, g)
	// 	}
	// }
	// return data, nil
}

func (r *PostgresGoalRepository) SaveGoal(goal model.Goal) error {
	_, err := r.DB.Exec(`
		INSERT INTO goals (user_id, title, description, target_date, status)
		VALUES ($1, $2, $3, $4, $5)
	`, goal.UserID, goal.Title, goal.Description, goal.TargetDate, goal.Status)
	return err
}

// GetGoalByIDは、特定の目標をIDで取得するメソッドです。
// このメソッドは、目標のIDを受け取り、データベースから目標の詳細を取得します。
// IDに紐づくstudy recordsも取得し、model.GoalDetailData型で返します。
// 目標の詳細には、目標自体の情報と、関連する学習記録のスライスが含まれます。
// エラーが発生した場合は、エラーを返します。
func (r *PostgresGoalRepository) GetGoalByID(id int) (model.GoalDetailData, error) {
	var goal model.GoalDetailData
	err := r.DB.QueryRow("SELECT id, user_id, title, description, target_date, status FROM goals WHERE id = $1", id).Scan(
		&goal.Goal.ID, &goal.Goal.UserID, &goal.Goal.Title, &goal.Goal.Description, &goal.Goal.TargetDate, &goal.Goal.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.GoalDetailData{}, nil // 目標が見つからない場合は空のGoalDetailDataを返す
		}
		return model.GoalDetailData{}, err // その他のエラーはそのまま返す
	}
	// 学習記録を取得
	rows, err := r.DB.Query("SELECT id, goal_id, content, duration_minutes, recorded_at FROM study_records WHERE goal_id = $1", id)
	if err != nil {
		return model.GoalDetailData{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var record model.StudyRecord
		if err := rows.Scan(&record.ID, &record.GoalID, &record.Content, &record.DurationMinutes, &record.RecordedAt); err != nil {
			return model.GoalDetailData{}, err
		}
		goal.StudyRecords = append(goal.StudyRecords, record)
	}
	return goal, nil
}

// SaveRecordは、学習記録をデータベースに保存するメソッドです。
// このメソッドは、学習記録のデータを受け取り、
// データベースに保存します。
// 学習記録の内容、学習時間、記録日時を含む必要があります。
// エラーが発生した場合は、エラーを返します。
func (r *PostgresGoalRepository) SaveRecord(record model.StudyRecord) error {
	_, err := r.DB.Exec(`
		INSERT INTO study_records (goal_id, content, duration_minutes, recorded_at)
		VALUES ($1, $2, $3, $4)
	`, record.GoalID, record.Content, record.DurationMinutes, record.RecordedAt)
	if err != nil {
		log.Printf("SaveRecord error: %v, goal_id=%v, content=%v, duration_minutes=%v, recorded_at=%v",
			err, record.GoalID, record.Content, record.DurationMinutes, record.RecordedAt)
	}
	return err
}

// UpdateGoalは、目標の情報を更新するメソッドです。
// このメソッドは、目標のIDを受け取り、
// 目標のタイトル、説明、ステータス、ターゲット日を更新します。
// 更新が成功した場合はnilを返し、
// 失敗した場合はエラーを返します。
func (r *PostgresGoalRepository) UpdateGoal(goal model.Goal) error {
	_, err := r.DB.Exec(`
		UPDATE goals
		SET title = $1, description = $2, status = $3, target_date = $4
		WHERE id = $5
	`, goal.Title, goal.Description, goal.Status, goal.TargetDate, goal.ID)
	if err != nil {
		log.Printf("UpdateGoal error: %v, goal_id=%v", err, goal.ID)
	}
	return err
}

// DeleteGoalは、目標を削除するメソッドです。
// このメソッドは、目標のIDを受け取り、
// データベースから目標を削除します。
// 削除が成功した場合はnilを返し、
// 失敗した場合はエラーを返します。
func (r *PostgresGoalRepository) DeleteGoal(id int) error {
	_, err := r.DB.Exec("DELETE FROM goals WHERE id = $1", id)
	if err != nil {
		log.Printf("DeleteGoal error: %v, goal_id=%v", err, id)
	}
	return err
}

// DeleteRecordは、学習記録を削除するメソッドです。
func (r *PostgresGoalRepository) DeleteRecord(id int) error {
	_, err := r.DB.Exec("DELETE FROM study_records WHERE id = $1", id)
	if err != nil {
		log.Printf("DeleteRecord error: %v, record_id=%v", err, id)
	}
	return err
}

func (r *PostgresGoalRepository) GetRecordByID(id int) (model.StudyRecord, error) {
	var record model.StudyRecord
	err := r.DB.QueryRow("SELECT id, goal_id, content, duration_minutes, recorded_at FROM study_records WHERE id = $1", id).Scan(
		&record.ID, &record.GoalID, &record.Content, &record.DurationMinutes, &record.RecordedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.StudyRecord{}, nil // 学習記録が見つからない場合は空のStudyRecordを返す
		}
		return model.StudyRecord{}, err // その他のエラーはそのまま返す
	}
	return record, nil
}

func (r *PostgresGoalRepository) UpdateRecord(record model.StudyRecord) error {
	_, err := r.DB.Exec(`
		UPDATE study_records
		SET content = $1, duration_minutes = $2, recorded_at = $3
		WHERE id = $4
	`, record.Content, record.DurationMinutes, record.RecordedAt, record.ID)
	if err != nil {
		log.Printf("UpdateRecord error: %v, record_id=%v", err, record.ID)
	}
	return err
}
