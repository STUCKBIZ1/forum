package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "bad request this method is not valid", 400)
		return
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
	posts, err := GetPosts()
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
