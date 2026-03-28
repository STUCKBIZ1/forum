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

	database.CreateTables(DB)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	fmt.Println("Server listen on http://localhost:8080/")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
