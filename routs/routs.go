package routs

import (
	"net/http"

	"forum/handlers"
)

func Routes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegiterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogOutHandler)
	http.HandleFunc("/api/{PC}/{id}/{action}", handlers.CLDhandlers)
	http.HandleFunc("/post/create", handlers.PostHandler)
}
