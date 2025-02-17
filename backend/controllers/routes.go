package controllers

import (
	"net/http"
	"personal-library/backend/models"
)

func SetupRoutes() {
	repo := &models.InMemoryBookRepository{}
	handler := NewBookHandler(repo)

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.AddBookHandler(w, r)
		case http.MethodGet:
			handler.GetAllBooksHandler(w, r)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetBookByIdHandler(w, r)
		case http.MethodPut:
			handler.EditBookByIdHandler(w, r)
		case http.MethodDelete:
			handler.DeleteBookByIdHandler(w, r)
		}
	})
}
