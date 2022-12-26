package main

import (
	"html/template"
	"net/http"

	"github.com/nathanieltruitt/first-go-web-app/views"
)

type Test struct {
	Thing string
}

func test(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./views/html/index.html")
	p := Test{Thing: "help"}
	tmpl.Execute(w, p)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		index := views.IndexPage{Name: "Nathan"}.RenderIndex()
		index.Page.Execute(w, index)
	default:
		http.Error(w, "Resource Not Found", http.StatusNotFound)
	}
}

func main() {
	nathanMux := http.NewServeMux()
	nathanMux.HandleFunc("/", handleFunc)
	nathanMux.HandleFunc("/test", test)
	nathanMux.HandleFunc("/login", views.Login)
	nathanMux.HandleFunc("/login_submit", views.LoginSubmit)
	server := http.Server{}
	server.Addr = ""
	server.Handler = nathanMux
	server.ListenAndServe()
}