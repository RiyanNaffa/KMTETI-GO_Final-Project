package main

import (
	"fmt"
	"net/http"

	"book-store/api"
)

func main() {
	h := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	// Test on root
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Testing..."))
	})
	// Handle book
	h.HandleFunc("/api/book", api.BookHandler)
	// Handle employee
	// h.HandleFunc("/api/employee",)

	fmt.Println("Running on Port:8080")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
