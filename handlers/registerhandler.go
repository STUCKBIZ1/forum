package handlers

import (
	"html/template"
	"net/http"
)

func RegiterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/register" {
		CreateUser(w, r)
	}
	if r.Method != http.MethodGet {
		// err
		return
	}
	if r.URL.Path != "/register" {
		// err
		return
	}	
	tmpl, err := template.ParseFiles("template/register.html")
	if err != nil {
		// err
		return
	}
	tmpl.Execute(w, nil)
}
