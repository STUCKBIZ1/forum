package handlers

import (
	"net/http"
	"html/template"
)

func RegiterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method	!= http.MethodGet{
		//err
		return
	}
	if r.URL.Path != "/register"{
		//err
		return
	}
	tmpl, err := template.ParseFiles("template/register.html")
	if err != nil{
		//err
		return
	}
	tmpl.Execute(w, nil)
}
