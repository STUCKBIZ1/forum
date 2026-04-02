package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func CreatPosthandler(w http.ResponseWriter, r *http.Request) {
	username, user_id, err := GetUserName(r)
	if err != nil {
		http.Error(w, "intrnal server error", 500)
		return
	}
	content := r.FormValue("content")
	category := r.FormValue("category")
	p := CreatCPLD{
		CreatPost : Post{
		ID:  user_id,
		Author:   username,
		Content:  content,
		Category: category,
	},
	}
	err = InsertingData(p, "post")
	if err != nil{
		log.Fatal("ERROR", err)
	}
	fmt.Println(username, user_id)
	http.Redirect(w, r, "/", 302)
}
