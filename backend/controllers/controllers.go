package controllers

import (
	"encoding/json"
	"net/http"
	"personal-library/backend/models"
	"strconv"
	"strings"
)

func (c *BookHandler) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newBook.Title == "" {
		http.Error(w, "Gagal menambahkan buku. Mohon isi nama buku", http.StatusBadRequest)
		return
	}

	if newBook.ReadPage > newBook.PageCount {
		http.Error(w, "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount", http.StatusBadRequest)
		return
	}

	book := c.Handler.AddBook(newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Buku berhasil ditambahkan",
		"data":    map[string]interface{}{"bookId": book.ID},
	})
}

func (c *BookHandler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	finished := query.Get("finished")

	filteredBooks := c.Handler.FilterBooks(name, finished)

	result := make([]map[string]interface{}, 0)
	for _, book := range filteredBooks {
		result = append(result, map[string]interface{}{
			"id":       book.ID,
			"title":    book.Title,
			"finished": book.Finished,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   map[string]interface{}{"books": result},
	})
}

func (c *BookHandler) GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/books/"))
	if err != nil {
		http.Error(w, "ID buku tidak valid", http.StatusBadRequest)
		return
	}
	book, found := c.Handler.GetBookByID(id)
	if !found {
		http.Error(w, "Buku tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   map[string]interface{}{"book": book},
	})
}

func (c *BookHandler) EditBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/books/"))
	if err != nil {
		http.Error(w, "ID buku tidak valid", http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedBook.ReadPage > updatedBook.PageCount {
		http.Error(w, "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount", http.StatusBadRequest)
		return
	}

	success := c.Handler.UpdateBook(id, updatedBook)

	if !success {
		http.Error(w, "Gagal memperbarui buku. Id tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Buku berhasil diperbarui",
	})
}

func (c *BookHandler) DeleteBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/books/"))
	if err != nil {
		http.Error(w, "ID buku tidak valid", http.StatusBadRequest)
		return
	}
	success := c.Handler.DeleteBook(id)
	if !success {
		http.Error(w, "Buku gagal dihapus. Id tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Buku berhasil dihapus",
	})
}
