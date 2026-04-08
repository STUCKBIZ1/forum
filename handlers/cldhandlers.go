package handlers

import (
	"net/http"
	"strconv"
)

func CLDhandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request", 400)
		return
	}
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "page not found", 404)
		return
	}
	action := r.PathValue("action")
	if !IsTrue(action) {
		http.Error(w, "page not found", 404)
		return
	}
	PC := r.PathValue("PC")
	if !IsExistPC(PC) {
		http.Error(w, "bad request", 400)
		return
	}
	if !SesIsExist(r) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	switch PC {
	case "post":
		switch action {
		case "comment":
			CommentHandler(w, r, n)
			return
		case "like":
			LikePostAndcommentHandler(w, r, n, "for post")
			return
		case "dislike":
			DislikePostAndcommentHandler(w, r, n, "for post")
			return
		default:
			http.Error(w, "page not found", 404)
			return
		}
	case "comment":
		switch action {
		case "like":
			LikePostAndcommentHandler(w, r, n, "for comment")
			return
		case "dislike":
			DislikePostAndcommentHandler(w, r, n, "for comment")
			return
		default:
			http.Error(w, "page not found", 404)
			return
		}
	}
}
