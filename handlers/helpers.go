package handlers

import (
	"database/sql"
)

func GetPosts(db *sql.DB) (any, error) {
	rows, err := db.Query("SELECT id, content, author, likes, dislikes FROM posts")
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
	return posts, nil
}
