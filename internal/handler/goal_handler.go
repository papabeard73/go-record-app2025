package handler

import (
	"go-record-app2025/internal/service"
	"html/template"
	"net/http"
)

type GoalHandler struct {
	GoalService *service.GoalService
}

func NewGoalHandler(goalService *service.GoalService) *GoalHandler {
	return &GoalHandler{GoalService: goalService}
}

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
