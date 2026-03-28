package handlers

import (
	"net/http"
	"strings"
)

func Commenthandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	if len(url) != 4 {
		// err
		return
	}
	// postid := url[2]
	action := url[3]
	if action != "comment" {
		// err
		return
	}
	_, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", 302)
		return
	}
}
