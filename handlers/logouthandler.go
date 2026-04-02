package handlers

import (
	"net/http"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "that not a valid", 422)
		return
	}
	if r.URL.Path != "/logout" {
		http.Error(w, "page not found", 422)
		return
	}
	session := GetToken(r)
	DeleteSession(session)
	http.Redirect(w, r, "/", 302)
}
