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
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
	);`

	_, err := db.Exec(postTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(commentTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tables created")
}
