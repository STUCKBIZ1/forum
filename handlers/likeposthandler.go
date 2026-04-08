package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func LikePostAndcommentHandler(w http.ResponseWriter, r *http.Request, ComOrPo_id int, category string) {
	username, _, err := GetUserName(r)
	if err != nil {
		log.Fatal("ERROR", err)
		return
	}
	var DI CreatCPLD
	var d Delete
	var categoryD string
	var categoryI string
	switch category {
	case "for post":
		d = Delete{
			Author:  username,
			Post_id: ComOrPo_id,
		}
		DI = CreatCPLD{
			L_DPostComment: L_DPostComment{
				Username: username,
				Post_id:  ComOrPo_id,
			},
		}
		categoryD = "from dislike post"
		categoryI = "from like post"
	case "for comment":
		d = Delete{
			Author:     username,
			Comment_id: ComOrPo_id,
		}
		DI = CreatCPLD{
			L_DPostComment: L_DPostComment{
				Username: username,
				Comment_id : ComOrPo_id,
			},
		}
		categoryD = "from dislike comment"
		categoryI = "from like comment"
	}
	_, err1 := GetData(username, categoryD, d)
	if err1 == nil {
		DeleteData(d, categoryD)
	}
	_, err = GetData(username, categoryI, d)
	if err != nil {
		err = InsertingData(DI, categoryI)
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
	} else {
		DeleteData(d, categoryI)
	}
	http.Redirect(w, r, "/", 302)
}
