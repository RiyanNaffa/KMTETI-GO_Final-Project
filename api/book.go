package api

import (
	"book-store/src/service"
	"encoding/json"
	"net/http"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		action := r.URL.Query().Get("action")

		switch action {
		case "DisplayAll":
			data, err := service.BookDisplayAll()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return

		case "Details":
			id := r.URL.Query().Get("id")
			if id == "" {
				http.Error(w, "No ID parsed.", http.StatusBadRequest)
				return
			}
			data, err := service.BookDetails(&id)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return
		}
	case http.MethodPut:
		// Change price & stock
		return
	case http.MethodPost:
		// Add a book
		w.Write([]byte("POST api/book"))
		return
	case http.MethodDelete:
		// Delete a book from database
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405: Method Not Allowed"))
		return
	}

}
