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
	posts, errr := GetPosts(DB)
	if errr != nil {
		// err
		return
	}
	temp, err := template.ParseFiles("template/home.html")
	if err != nil {
		// error
		return
	}
	fmt.Println(posts)

	temp.Execute(w, posts)
}
