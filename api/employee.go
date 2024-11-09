package api

import (
	"book-store/src/service"
	"encoding/json"
	"net/http"
)

// A function for handling various requests regarding the employee collection.
//
// Available HTTP methods are GET, POST and DELETE. EmployeeHandler is capable of displaying all employees'
// information in the database, insert an employee data into the database and delete an employee from the
// database.
func EmployeeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		action := r.URL.Query().Get("action")

		switch action {
		case "display":
			data, err := service.EmployeeDisplayAll()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&data)
			return

		default:
			http.Error(w, "query not accepted", http.StatusUnprocessableEntity)
			return
		}

	case http.MethodPost:
		action := r.URL.Query().Get("action")

		switch action {
		case "add":
			response, err := service.EmployeeAdd(r.Body)
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
				"message": "Employee successfully added.",
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

			response, err := service.EmployeeDelete(&id)

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
				"message":  "Employee deleted successfully.",
				"employee": &response,
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
