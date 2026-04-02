package handlers

import (
	"log"
	"net/http"
)

func CommentHandler(w http.ResponseWriter, r *http.Request, post_id int) {
	username, _, err := GetUserName(r)
	if err != nil {
		http.Error(w, "you give the empty comment", 404)
	}
	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "you give the empty comment", 404)
		return
	}
	C := CreatCPLD{
		CreatComment : Comment{
			ID: post_id,
			Author:  username,
			Content: content,
		},
	}
	err = InsertingData(C, "comment")
	if err != nil{
		log.Fatal("ERROR", err)
		return
	}
	http.Redirect(w, r, "/", 302)
	
}
