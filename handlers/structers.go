package handlers

import (
	"database/sql"
)

var DB *sql.DB

type Post struct {
	ID       int
	Title string
	Author   string
	Content  string
	Comments []Comment
    Category string
	Like     int
	Dislike  int
}

type CreatCPLD struct {
	CreatPost    Post
	CreatComment Comment
	L_DPostComment L_DPostComment
}
type L_DPostComment struct{
	ID int
	Post_id int
	Comment_id int
	Username string
}
type Posts struct {
	AllPosts []Post
	LoggedIn bool
	Username string
	General bool
	Webdeveloper bool
	Reports bool
	All bool
}
type Comment struct {
	ID      int
	PostID  int
	Author  string
	Content string
	Like    int
	Dislike int
}
type Delete struct{
	Author string
	Post_id int
	Comment_id int
	session string
}
type LikedDisliked struct{
	liked bool
	disliked bool
}