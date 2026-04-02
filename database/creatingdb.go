package database

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTables(db *sql.DB) {
	postTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		likes INTEGER DEFAULT 0,
		dislikes INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	commentTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		author TEXT NOT NULL,
		content TEXT NOT NULL,
		likes INTEGER DEFAULT 0,
		dislikes INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	userTable := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	session_user := `
	CREATE TABLE IF NOT EXISTS session_user(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER,
	session_token TEXT UNIQUE NOT NULL,
	username TEXT UNIQUE NOT NULL
	)
	`
	// likeTable := `	
	// CREATE TABLE IF NOT EXISTS like(
	// 	post_id INTEGER PRIMARY KEY,
	// 	user_id INTEGER PRIMARY KEY,
	// )
	// `
	// dislikeTable := `
	// 	CREATE TABLE IF NOT EXISTS like(
	// 	post_id INTEGER PRIMARY KEY,
	// 	user_id INTEGER PRIMARY KEY,
	// (
	// `
	_, err := db.Exec(session_user)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(userTable)
	if err != nil {
		log.Fatal(err)
	}
	// _, err = db.Exec(likeTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = db.Exec(dislikeTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	_, err = db.Exec(postTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(commentTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tables created")
}
