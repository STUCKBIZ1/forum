package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r)
	if !SesIsExist(r){
		http.Redirect(w, r, "/login", 302)
		return
	}
	switch r.Method {

	case http.MethodPost:
		CreatPosthandler(w, r)
		return

	case http.MethodGet:
		tmpl, err := template.ParseFiles("template/post.html")
		if err != nil {
			http.Error(w, "template not found", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
