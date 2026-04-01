package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"forum/database"
	"forum/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	DB, err := sql.Open("sqlite3", "./form.db")
	if err != nil {
		fmt.Println("DB open error:", err)
		return
	}
	defer DB.Close()

	if err := DB.Ping(); err != nil {
		fmt.Println("DB connection error:", err)
		return
	}

	handlers.DB = DB

	database.CreateTables(DB)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegiterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogOutHandler)
	http.HandleFunc("/post/", handlers.CLDPhandlers)
	fmt.Println("Server listen on http://localhost:8888/")
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
