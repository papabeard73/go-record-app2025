package handler

import (
	"go-record-app2025/internal/model"
	"go-record-app2025/internal/service"
	"html/template"
	"log"
	"net/http"
)

// GoalHandlerは、目標に関連するHTTPリクエストを処理するためのハンドラーです。
// このハンドラーは、目標の一覧を表示するための機能を提供します。
// 目標のデータは、GoalServiceを通じて取得されます
type GoalHandler struct {
	GoalService *service.GoalService
}

// インスタンスを生成する関数
// GoalHandlerのインスタンスを生成するための関数
func NewGoalHandler(goalService *service.GoalService) *GoalHandler {
	return &GoalHandler{GoalService: goalService}
}

// ListGoalsは、目標の一覧を表示するハンドラー関数です。
// この関数は、HTTPリクエストを受け取り、目標の一覧を取得して、HTMLテンプレートにレンダリングします。
// 目標のステータは、GoalServiceを通じて取得されます。
// 取得した目標データは、HTMLテンプレートに渡され、最終的にHTTPレスポンスとしてクライアントに返されます。
func (h *GoalHandler) ListGoals(w http.ResponseWriter, r *http.Request) {
	goals, err := h.GoalService.GetGoals()
	if err != nil {
		http.Error(w, "データ取得失敗", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.New("layout.html").Funcs(template.FuncMap{
		"eq": func(a, b string) bool { return a == b },
	}).ParseFiles("templates/layout.html", "templates/goal_list.html"))

	tmpl.ExecuteTemplate(w, "layout.html", goals)
}

func (h *GoalHandler) AddNewGoals(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.New("layout.html").Funcs(template.FuncMap{
			"eq": func(a, b string) bool { return a == b },
		}).ParseFiles("templates/layout.html", "templates/goal_new.html"))
		err := tmpl.ExecuteTemplate(w, "layout.html", nil)
		if err != nil {
			http.Error(w, "テンプレート描画エラー", http.StatusInternalServerError)
			log.Println("execute error:", err)
		}
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "フォーム解析エラー", http.StatusBadRequest)
			return
		}

		goal := model.Goal{
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
			TargetDate:  r.FormValue("target_date"),
			Status:      "NotStarted",
			UserID:      1, // 今は仮に固定
		}

		err := h.GoalService.CreateGoal(goal)
		if err != nil {
			http.Error(w, "登録失敗", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
