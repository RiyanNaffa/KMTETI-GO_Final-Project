package api

import (
	"net/http"
	// "book-store/src/service"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("GET api/book"))
	case "POST":
		w.Write([]byte("POST api/book"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
