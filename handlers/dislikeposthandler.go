package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func DislikePostHandler(w http.ResponseWriter, r *http.Request, post_id int) {
	username, _, err := GetUserName(r)
	if err != nil {
		log.Fatal("ERROR", err)
		return
	}
	d := Delete{
		Author:  username,
		Post_id: post_id,
	}
	_, err1 := GetData(username, "form like", d)
	if err1 != nil {
		DeleteData(d, "from dislike")
	}
	_, err = GetData(username, "from dislike", d)
	fmt.Println(err)
	if err != nil {
		d := CreatCPLD{
			LikePost: LikePost{
				Username: username,
				Post_id:  post_id,
			},
		}
		err = InsertingData(d, "dislikepost")
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
	}
	http.Redirect(w, r, "/", 302)
}
