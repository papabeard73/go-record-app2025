package handler

import (
	"go-record-app2025/internal/model"
	"go-record-app2025/internal/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
			TargetDate:  r.FormValue("target_date"),
			Status:      r.FormValue("status"),
			Description: r.FormValue("description"),
			UserID:      1, // 今は仮に固定
		}

		// デバッグ用
		// log.Println("title:", r.FormValue("title"))
		// log.Println("description:", r.FormValue("description"))
		// log.Println("target_date:", r.FormValue("target_date"))
		// log.Println("status:", r.FormValue("status"))
		// log.Printf("Received goal: %+v\n", goal)

		err := h.GoalService.CreateGoal(goal)
		if err != nil {
			log.Println("CreateGoal error:", err)
			http.Error(w, "登録失敗", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// DetailGoalsは、特定の目標の詳細を表示するハンドラー関数です。
// この関数は、目標のIDをURLパラメータから取得し、その目標の詳細をGoalServiceを通じて取得します。
// 取得した目標データは、HTMLテンプレートに渡され、最終的にHTTPレスポンスとしてクライアントに返されます。
func (h *GoalHandler) DetailGoals(w http.ResponseWriter, r *http.Request) {
	goalID := r.URL.Query().Get("id")
	if goalID == "" {
		http.Error(w, "目標IDが指定されていません", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(goalID)
	if err != nil {
		http.Error(w, "無効な目標ID", http.StatusBadRequest)
		return
	}
	goal, err := h.GoalService.DetailGoals(id)
	if err != nil {
		http.Error(w, "目標の詳細取得失敗", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.New("layout.html").Funcs(template.FuncMap{
		"eq": func(a, b string) bool { return a == b },
	}).ParseFiles("templates/layout.html", "templates/goal_detail.html"))
	err = tmpl.ExecuteTemplate(w, "layout.html", goal)
	if err != nil {
		http.Error(w, "テンプレート描画エラー", http.StatusInternalServerError)
		log.Println("execute error:", err)
		return
	}
}
