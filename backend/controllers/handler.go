package controllers

import (
	"personal-library/backend/models"
)

type BookHandler struct {
	Handler models.BookRepository
}

func NewBookHandler(repo models.BookRepository) *BookHandler {
	return &BookHandler{Handler: repo}
}

func (h *BookHandler) AddBook(book models.Book) models.Book {
	return h.Handler.AddBook(book)
}

func (h *BookHandler) GetAllBooks() []models.Book {
	return h.Handler.GetAllBooks()
}

func (h *BookHandler) GetBookByID(id int) (*models.Book, bool) {
	return h.Handler.GetBookByID(id)
}

func (h *BookHandler) UpdateBook(id int, updatedBook models.Book) bool {
	return h.Handler.UpdateBook(id, updatedBook)
}

func (h *BookHandler) DeleteBook(id int) bool {
	return h.Handler.DeleteBook(id)
}

func (h *BookHandler) FilterBooks(name, finished string) []models.Book {
	return h.Handler.FilterBooks(name, finished)
}
