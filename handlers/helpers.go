package handlers

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetPosts(slice []string) ([]Post, error) {
	rows, err := DB.Query("SELECT id, title, content, author, likes, dislikes, category FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var p Post
		rows.Scan(&p.ID, &p.Title, &p.Content, &p.Author, &p.Like, &p.Dislike, &p.Category)
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
		if slices.Contains(slice, p.Category) && len(slice) != 0 {
			posts = append(posts, p)
			continue
		}else if len(slice) == 0 {
			posts = append(posts, p)
		}else{
			continue
		}
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
	d := Delete{
		Author: username,
	}
		DeleteData(d, "from session_user")
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
	case "from session_user":
		_, err = DB.Exec("DELETE FROM session_user WHERE username = ?", d.Author)
	case "from session":
		_, err = DB.Exec("DELETE FROM session_user WHERE session_token = ?", d.session)
	case "from dislike post":
		_, err = DB.Exec("DELETE FROM dislikepost WHERE post_id = ? AND Author = ?", d.Post_id, d.Author)
		if err == nil {
			_, err = DB.Exec(`
	UPDATE posts 
	SET dislikes = dislikes - 1 
	WHERE id = ? AND dislikes > 0
`, d.Post_id)
		}
	case "from dislike comment":
		_, err = DB.Exec("DELETE FROM dislikecomment WHERE comment_id = ? AND Author = ?", d.Comment_id, d.Author)
		if err == nil {
			_, err = DB.Exec(`
	UPDATE comments 
	SET dislikes = dislikes - 1 
	WHERE id = ? AND dislikes > 0
`, d.Comment_id)
		}
	case "from like post":
		_, err = DB.Exec("DELETE FROM likepost WHERE post_id = ? AND Author = ?", d.Post_id, d.Author)

		if err == nil {
			_, err = DB.Exec(`
	UPDATE posts 
	SET likes = likes - 1 
	WHERE id = ? AND likes > 0
`, d.Post_id)
		}

	case "from like comment":
		_, err = DB.Exec("DELETE FROM likecomment WHERE comment_id = ? AND Author = ?", d.Comment_id, d.Author)

		if err == nil {
			_, err = DB.Exec(`
	UPDATE comments 
	SET likes = likes - 1 
	WHERE id = ? AND likes > 0
`, d.Comment_id)
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

func IsExistPC(s string) bool {
	switch s {
	case "post", "comment":
		return true
	}
	return false
}

func InsertingData(s CreatCPLD, category string) error {
	var err error
	switch category {
	case "post":
		_, err = DB.Exec("INSERT INTO posts (user_id, title, content, author, category) VALUES (?, ?, ?, ?, ?)", s.CreatPost.ID, s.CreatPost.Title, s.CreatPost.Content, s.CreatPost.Author, s.CreatPost.Category)
		if err != nil {
			return err
		}
	case "comment":
		_, err = DB.Exec("INSERT INTO comments (post_id, author, content) VALUES (?, ?, ?)", s.CreatComment.ID, s.CreatComment.Author, s.CreatComment.Content)
	case "from like post":
		_, err = DB.Exec("INSERT INTO likepost (post_id, Author) VALUES (?, ?)", s.L_DPostComment.Post_id, s.L_DPostComment.Username)
		if err == nil {
			_, err = DB.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", s.L_DPostComment.Post_id)
		}
	case "from like comment":
		_, err = DB.Exec("INSERT INTO likecomment (comment_id, Author) VALUES (?, ?)", s.L_DPostComment.Comment_id, s.L_DPostComment.Username)
		if err == nil {
			_, err = DB.Exec("UPDATE comments SET likes = likes + 1 WHERE id = ?", s.L_DPostComment.Comment_id)
		}
	case "from dislike post":
		_, err = DB.Exec("INSERT INTO dislikepost (post_id, Author) VALUES (?, ?)", s.L_DPostComment.Post_id, s.L_DPostComment.Username)
		if err == nil {
			_, err = DB.Exec("UPDATE posts SET dislikes = dislikes + 1 WHERE id = ?", s.L_DPostComment.Post_id)
		}
	case "from dislike comment":
		_, err = DB.Exec("INSERT INTO dislikecomment (comment_id, Author) VALUES (?, ?)", s.L_DPostComment.Comment_id, s.L_DPostComment.Username)
		if err == nil {
			_, err = DB.Exec("UPDATE comments SET dislikes = dislikes + 1 WHERE id = ?", s.L_DPostComment.Comment_id)
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
	case "from like post":
		err = DB.QueryRow(
			"SELECT post_id, Author FROM likepost WHERE post_id = ? AND Author = ?",
			d.Post_id, d.Author,
		).Scan(&post_id, &author)
	case "from like comment":
		err = DB.QueryRow(
			"SELECT comment_id, Author FROM likecomment WHERE comment_id = ? AND Author = ?",
			d.Comment_id, d.Author,
		).Scan(&post_id, &author)

	case "from dislike post":
		err = DB.QueryRow(
			"SELECT post_id, Author FROM dislikepost WHERE post_id = ? AND Author = ?",
			d.Post_id, d.Author,
		).Scan(&post_id, &author)
	case "from dislike comment":
		err = DB.QueryRow(
			"SELECT comment_id, Author FROM dislikecomment WHERE comment_id = ? AND Author = ?",
			d.Comment_id, d.Author,
		).Scan(&post_id, &author)
	}
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
