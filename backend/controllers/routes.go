package controllers

import (
	"net/http"
	"personal-library/backend/models"
)

func SetupRoutes() {
	repo := &models.InMemoryBookRepository{}
	handler := models.NewBookHandler(repo)
	controller := NewBookController(handler)

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controller.AddBookHandler(w, r)
		case http.MethodGet:
			controller.GetAllBooksHandler(w, r)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetBookByIdHandler(w, r)
		case http.MethodPut:
			controller.EditBookByIdHandler(w, r)
		case http.MethodDelete:
			controller.DeleteBookByIdHandler(w, r)
		}
	})
}
