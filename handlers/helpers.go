package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetPosts() ([]Post, error) {
	rows, err := DB.Query("SELECT id, content, author, likes, dislikes FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Content, &p.Author, &p.Like, &p.Dislike)
		commentrows, err := DB.Query("SELECT id, post_id, author, content, likes, dislikes FROM comments WHERE post_id = ? ORDER BY created_at DESC", p.ID)
		if err != nil {
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
	user_id, err := GetData(username, "from user", Delete{})
	if err != nil {
		log.Fatal("ERROR", err)
		return
	}
	token := uuid.New().String()
	query := `INSERT INTO session_user (user_id, session_token, username) VALUES (?, ?, ?)`

	_, err = DB.Exec(query, user_id, token, username)
	if err != nil {
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
	var s string
	err := DB.QueryRow("SELECT session_token FROM session_user WHERE session_token = ?", session).Scan(&s)
	if err != nil {
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

func GetUserName(r *http.Request) (string, int, error) {
	token := GetToken(r)
	var u string
	var user_id int
	err := DB.QueryRow("SELECT username, user_id FROM session_user WHERE session_token = ?", token).Scan(&u, &user_id)
	if err != nil {
		// err
		return "", 0, err
	}
	return u, user_id, nil
}

func DeleteData(d Delete, category string) {
	var err error
	switch category {
	case "from session":
		_, err = DB.Exec("DELETE FROM session_user WHERE session_token = ?", d.session)
	case "from dislike":
		_, err = DB.Exec("DELETE FROM dislike WHERE post_id = ? AND Author = ?", d.Post_id, d.Author)
		if err == nil {
			_, err = DB.Exec(`
	UPDATE posts 
	SET dislikes = dislikes - 1 
	WHERE id = ? AND dislikes > 0
`, d.Post_id)
		}
	case "from like":
		_, err = DB.Exec("DELETE FROM like WHERE post_id = ? AND Author = ?", d.Post_id, d.Author)

		if err == nil {
			_, err = DB.Exec(`
	UPDATE posts 
	SET likes = likes - 1 
	WHERE id = ? AND likes > 0
`, d.Post_id)
		}
	}

	if err != nil {
		log.Fatal("Failed to delete row:", err)
	}
}

func IsTrue(s string) bool {
	switch s {
	case "comment", "like", "dislike":
		return true
	}
	return false
}

func InsertingData(s CreatCPLD, category string) error {
	var err error
	switch category {
	case "post":
		_, err = DB.Exec("INSERT INTO posts (user_id, content, author) VALUES (?, ?, ?)", s.CreatPost.ID, s.CreatPost.Content, s.CreatPost.Author)
		if err != nil {
			return err
		}
	case "comment":
		_, err = DB.Exec("INSERT INTO comments (post_id, author, content) VALUES (?, ?, ?)", s.CreatComment.ID, s.CreatComment.Author, s.CreatComment.Content)
	case "likepost":
		_, err = DB.Exec("INSERT INTO like (post_id, Author) VALUES (?, ?)", s.LikePost.Post_id, s.LikePost.Username)
		if err == nil {
			_, err = DB.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", s.LikePost.Post_id)

		}
	case "dislikepost":
		_, err = DB.Exec("INSERT INTO dislike (post_id, Author) VALUES (?, ?)", s.LikePost.Post_id, s.LikePost.Username)
		if err == nil {
			_, err = DB.Exec("UPDATE posts SET dislikes = dislikes + 1 WHERE id = ?", s.LikePost.Post_id)
		}
	}

	if err != nil {
		return err
	}
	return err
}

func GetData(username string, category string, d Delete) (int, error) {
	var err error
	var author string
	var post_id int
	var user_id int
	switch category {
	case "from user":
		err = DB.QueryRow("SELECT id FROM user WHERE username = ?", username).Scan(&user_id)
	case "from like":
		err = DB.QueryRow(
			"SELECT post_id, Author FROM like WHERE post_id = ? AND Author = ?",
			d.Post_id, d.Author,
		).Scan(&post_id, &author)

	case "from dislike":
		err = DB.QueryRow(
			"SELECT post_id, Author FROM dislike WHERE post_id = ? AND Author = ?",
			d.Post_id, d.Author,
		).Scan(&post_id, &author)
	}
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
