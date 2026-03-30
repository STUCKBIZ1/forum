package handlers

import (
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/creat-account" {
		// err
		return
	}
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println(username, email, password)
	fmt.Fprintln(w, "hwllo")
}
