package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var slice []string
	if r.Method == http.MethodPost {
		if gene := r.FormValue("General"); gene == "General" {
			slice = append(slice, gene)
		}
		if webdeveloper := r.FormValue("Webdeveloper"); webdeveloper == "Webdeveloper" {
			slice = append(slice, webdeveloper)
		}
		if rep := r.FormValue("Reports"); rep == "Reports" {
			slice = append(slice, rep)
		}
	}
	if r.URL.Path != "/" {
		http.Error(w, "not found", 404)
		return
	}
	var username string
	checkse := SesIsExist(r)
	if checkse {
		username, _, _ = GetUserName(r)
	}
	posts, err := GetPosts(slice)
	p := Posts{
		AllPosts: posts,
		LoggedIn: checkse,
		Username: username,
	}
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	temp, err := template.ParseFiles("template/home.html")
	if err != nil {
		// error
		return
	}

	temp.Execute(w, p)
}
