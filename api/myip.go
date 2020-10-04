package handler

import (
	"fmt"
	"net/http"
)

//MyIP handle request for ...
func MyIP(w http.ResponseWriter, r *http.Request) {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		fmt.Fprintf(w, forwarded)
		return
	}
	fmt.Fprintf(w, r.RemoteAddr)
}
