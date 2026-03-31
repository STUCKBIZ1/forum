package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"database/sql"
	"fmt"
	"net/http"
)

func GetPosts(db *sql.DB) (any, error) {
	rows, err := db.QueryRow("SELECT id, content, author, likes, dislikes FROM posts")
	if err != nil {
		// err
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Content, &p.Author, &p.Like, &p.Dislike)
		commentrows, err := db.Query("SELECT id, post_id, author, content, likes, dislikes FROM comments WHERE post_it = ?", p.ID)
		if err != nil {
			// err
			return nil, err
		}
		defer commentrows.Close()
		for commentrows.Next() {
			var c Comment
			commentrows.Scan(&c.ID, &c.PostID, &c.Author, &c.Content, &c.Like, &c.Dislike)
			p.Comments = append(p.Comments, c)
		}
		posts = append(posts, p)
	}
	posts = append(posts, Post{ID: 1, Author: "Ahmed", Content: "dfs World!", Like: 10, Dislike: 2})
	posts = append(posts, Post{ID: 1, Author: "ali", Content: "Helererereelo Wodfdfrld!", Like: 11, Dislike: 2})
	posts = append(posts, Post{ID: 1, Author: "Ahed", Content: "Hello World!", Like: 10, Dislike: 2})
	posts = append(posts, Post{ID: 1, Author: "moh", Content: "Hellererero Wdfdorld!", Like: 10, Dislike: 2})
	posts = append(posts, Post{ID: 1, Author: "ueu", Content: "Hello World!", Like: 12340, Dislike: 2})
	posts = append(posts, Post{ID: 1, Author: "diier", Content: "Hellodfdfeferer World!", Like: 11, Dislike: 2})

	return posts, nil
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	var dbpassowrd 
	row := DB.QueryRow("SELECT password FROME user WHERE username = ?", username)
	err := row.Scan(&dbpassowrd)

	//  cookie := &http.Cookie{
	// 	Name : "session_token",
	// 	Value : 
		
	//  }
	//  http.SetCookie(w, cookie)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	query := `INSERT INTO user(username, email, password) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, username, email, password)
	if err != nil {
		// err
		fmt.Fprintln(w, err)
		return
	}
	http.Redirect(w, r, "/login", 302)
}
