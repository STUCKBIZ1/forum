package handlers

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/login" {
		SignIn(w, r)
		return
	}
	if r.Method != http.MethodGet {
		// err
		return
	}
	if r.URL.Path != "/login" {
		// err
		return
	}
	tmpl, err := template.ParseFiles("template/login.html")
	if err != nil {
		// err
		return
	}

	tmpl.Execute(w, nil)
}
