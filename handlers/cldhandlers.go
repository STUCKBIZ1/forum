	package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func CLDPhandlers(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	_, err := strconv.Atoi(url[2])
	// var action string
	// var postid string
	if len(url) == 4 && err == nil && istrue(url[3])&& url[0] == "" && r.Method == http.MethodPost{
	// action = url[3]
	// postid = url[2]
	}else if len(url) == 3 && istrue(url[2])&& url[0] == ""&& r.Method == http.MethodGet{
		// action = url[2]
	}else{
		//err
		return
	}
	_, err = r.Cookie("session_token")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", 302)
		return
	}
}
func istrue(s string) bool{
	switch s{
	case "comment", "like", "dislike", "create":
		return true
	}
	return false
}
