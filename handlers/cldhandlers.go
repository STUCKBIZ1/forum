package handlers

import (
	"net/http"
	"strconv"
)

func CLDPhandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost{
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
	yes := SesIsExist(r)
	if !yes {
		http.Redirect(w, r, "/login", 302)
		return
	}
	switch action {
	case "comment":
		CommentHandler(w, r, n)
		return
	case "like":
		LikePostHandler(w, r, n)
		return
	case "dislike":
		DislikePostHandler(w, r, n)
		return
	default:
		http.Error(w, "page not found", 404)
		return
	}
}
