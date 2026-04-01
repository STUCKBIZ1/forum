package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// error
		return
	}
	if r.URL.Path != "/" {
		// error
		return
	}
	var username string 
	checkse := SesIsExist(r)
	if checkse{
		username = GetUserName(r)
	}
	fmt.Println(checkse)
	posts, err := GetPosts()
	p := Posts{
		AllPosts: posts,
		LoggedIn: checkse,
		Username: username,
	}
	if err != nil {
		// err
		return
	}
	temp, err := template.ParseFiles("template/home.html")
	if err != nil {
		// error
		return
	}

	temp.Execute(w, p)
}
