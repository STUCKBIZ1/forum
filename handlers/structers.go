package handlers
import (
    "database/sql"
)
var DB *sql.DB

type Post struct {
    ID       int
    Author   string
    Content  string
    Comments []Comment
    Like     int
    Dislike  int
}
type Posts struct{
    AllPosts []Post
    LoggedIn bool
    Username string
}
type Comment struct {
    ID      int
    PostID  int
    Author  string
    Content string
    Like    int
    Dislike int
}