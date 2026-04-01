package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetPosts() ([]Post, error) {
	rows, err := DB.Query("SELECT id, content, author, likes, dislikes FROM posts")
	if err != nil {
		// err
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Content, &p.Author, &p.Like, &p.Dislike)
		commentrows, err := DB.Query("SELECT id, post_id, author, content, likes, dislikes FROM comments WHERE post_it = ?", p.ID)
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
	password := r.FormValue("password")
	username := r.FormValue("username")

	var dbpassowrd string
	err := DB.QueryRow("SELECT password FROM user WHERE username = ?", username).Scan(&dbpassowrd)
	if err != nil {
		http.Error(w, "User not found", 400)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbpassowrd), []byte(password))
	if err != nil {
		http.Error(w, "Wrong password", 400)
		return
	}

	token := uuid.New().String()
	query := `INSERT INTO session_user (session_token, username) VALUES (?, ?)`

	_, err = DB.Exec(query, token, username)
	if err != nil {
		fmt.Println("ERROR", err)
		http.Error(w, "ERROR", 422)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	dbpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// err
		return
	}
	query := `INSERT INTO user(username, email, password) VALUES (?, ?, ?)`
	_, err = DB.Exec(query, username, email, string(dbpassword))
	if err != nil {
		// err
		fmt.Fprintln(w, err)
		return
	}
	http.Redirect(w, r, "/login", 302)
}

func SesIsExist(r *http.Request) bool {
	session := GetToken(r)
	fmt.Println(session)
	var s string
	err := DB.QueryRow("SELECT session_token FROM session_user WHERE session_token = ?", session).Scan(&s)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetToken(r *http.Request) string {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// err
		return ""
	}
	token := cookie.Value
	return token
}

func GetUserName(r *http.Request) string {
	token := GetToken(r)
	var u string
	err := DB.QueryRow("SELECT  username FROM session_user WHERE session_token = ?", token).Scan(&u)
	if err != nil {
		// err
		return ""
	}
	return u
}
func DeleteSession(session string){
	_, err := DB.Exec("DELETE FROM session_user WHERE session_token = ?", session)
	if err != nil{
		log.Fatal("Failed to delete row:", err)
	}
	return
}
