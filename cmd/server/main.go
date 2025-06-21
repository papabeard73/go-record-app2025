package main

import (
	"html/template"
	"log"
	"net/http"
)

type Goal struct {
	ID          int
	UserId      int
	Title       string
	Description string
	TargetDate  string
	Status      string // "NotStarted", "ActiveGoals", "CompletedGoals"
}

func main() {
	// 静的ファイル (例: /static/images/xxx.svg)
	// staticというディレクトリをルートにしたファイルサーバーを作る。つまり、staticフォルダの中身をWeb経由で見せる準備
	fs := http.FileServer(http.Dir("static"))
	// /static/ というパスでアクセスが来たら、staticディレクトリの中身を見せる
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	goals := []Goal{
		{ID: 1, Title: "Javaの学習（Paiza、Youtube、デイトラ）", TargetDate: "2025/08/15", Status: "NotStarted"},
		{ID: 2, Title: "Go言語でのアプリ開発", TargetDate: "2025/08/20", Status: "ActiveGoals"},
		{ID: 3, Title: "TOEIC受験", TargetDate: "2025/01/10", Status: "CompletedGoals"},
	}

	// "/" にアクセスが来たら goal_list.html を返す
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// layout.htmlという名前のテンプレートを新しく作る。そのあとに .Funcs(...) とつなげ、「テンプレートの中で使える関数」を追加
		tmpl := template.New("layout.html").Funcs(template.FuncMap{
			// テンプレートで使える関数を定義
			"eq": func(a, b string) bool { return a == b },
		})
		// 複数テンプレートを読み込み
		var err error
		tmpl, err = tmpl.ParseFiles(
			"templates/layout.html",
			"templates/goal_list.html",
		)
		if err != nil {
			http.Error(w, "テンプレート読み込みエラー", http.StatusInternalServerError)
			log.Println("template error:", err)
			return
		}
		// goal_list.htmlを描画する
		// テンプレートに渡すデータとして goals を渡す
		err = tmpl.ExecuteTemplate(w, "layout.html", goals)
		if err != nil {
			http.Error(w, "テンプレート描画エラー", http.StatusInternalServerError)
			log.Println("execute error:", err)
		}
	})

	log.Println("🚀 http://localhost:8080 でサーバー起動中")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバー起動失敗:", err)
	}
}
