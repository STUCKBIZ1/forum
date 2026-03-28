package handlers

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "WELCOMME TO LOGIN PAGE")
}
