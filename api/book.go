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

		case "display":
			data, err := service.BookDisplayAll()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
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
				switch err.Error() {
				case "internal server error":
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				case "not found":
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&data)

			return

		default:
			http.Error(w, "Query Not Accepted.", http.StatusUnprocessableEntity)
			return
		}

	case http.MethodPut:
		action := r.URL.Query().Get("action")

		switch action {
		case "update":
			response, err := service.BookUpdate(r.Body)
			if err != nil {
				switch err.Error() {
				case "bad request":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "unprocessable entity":
					http.Error(w, err.Error(), http.StatusUnprocessableEntity)
					return
				case "internal server error":
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			message := "Book successfully updated."

			if response.MatchedCount == 0 {
				message = "Book not found."
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":  message,
				"modified": response.ModifiedCount,
			})

		default:
			http.Error(w, "Query Not Accepted.", http.StatusUnprocessableEntity)
			return
		}
		return

	case http.MethodPost:
		action := r.URL.Query().Get("action")

		switch action {

		case "add":
			response, err := service.BookAdd(r.Body)
			if err != nil {
				switch err.Error() {
				case "internal server error":
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				case "bad request":
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				case "unprocessable entity":
					http.Error(w, err.Error(), http.StatusUnprocessableEntity)
					return
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Book successfully added.",
				"_id":     response.InsertedID,
			})

			return

		default:
			http.Error(w, "query not accepted", http.StatusUnprocessableEntity)
			return
		}

	case http.MethodDelete:
		action := r.URL.Query().Get("action")

		switch action {

		case "delete":
			id := r.URL.Query().Get("id")

			if id == "" {
				http.Error(w, "No ID parsed.", http.StatusBadRequest)
				return
			}

			response, err := service.BookDelete(&id)

			if err != nil {
				switch err.Error() {
				case "internal server error":
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				case "not found":
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Book deleted successfully.",
				"book":    &response,
			})

			return
		default:
			http.Error(w, "query not accepted", http.StatusUnprocessableEntity)
			return
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
