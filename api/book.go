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

		case "displayAll":
			data, err := service.BookDisplayAll()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&data)
			return

		case "details":
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

		default:
			http.Error(w, "Query Not Accepted.", http.StatusUnprocessableEntity)
			return
		}

	case http.MethodPut:
		// Change price & stock
		return

	case http.MethodPost:
		if r.URL.Query().Get("action") != "add" {
			http.Error(w, "Query Not Accepted.", http.StatusUnprocessableEntity)
			return
		}

		response, err := service.BookAdd(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Location", fmt.Sprint("/api/book&add"))
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Book successfully added.",
			"id":      response.InsertedID,
		})

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
