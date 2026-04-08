package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"forum/database"
	"forum/handlers"
	"forum/routs"

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
	routs.Routes()

	database.CreateTables(DB)
	fmt.Println("Server listen on http://localhost:8888/")
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
