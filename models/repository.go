package models

import (
	"strings"
	"time"
)

type Book struct {
	ID         int
	Title      string
	Year       int
	Author     string
	PageCount  int
	ReadPage   int
	Finished   bool
	Reading    bool
	InsertedAt string
	UpdatedAt  string
}

var books []Book
var lastID int = 0

type BookRepository interface {
	AddBook(book Book) Book
	GetAllBooks() []Book
	GetBookByID(id int) (*Book, bool)
	UpdateBook(id int, updatedBook Book) bool
	DeleteBook(id int) bool
	FilterBooks(name, reading, finished string) []Book
}

type InMemoryBookRepository struct{}

// UpdateBook implements BookRepository.
func (r *InMemoryBookRepository) UpdateBook(id int, updatedBook Book) bool {
	panic("unimplemented")
}

func (r *InMemoryBookRepository) AddBook(book Book) Book {
	lastID++
	book.ID = lastID
	book.InsertedAt = time.Now().Format(time.RFC3339)
	book.UpdatedAt = book.InsertedAt
	book.Finished = book.PageCount == book.ReadPage

	books = append(books, book)
	return book
}

func (r *InMemoryBookRepository) GetAllBooks() []Book {
	return books
}

func (r *InMemoryBookRepository) GetBookByID(id int) (*Book, bool) {
	for _, book := range books {
		if book.ID == id {
			return &book, true
		}
	}
	return nil, false
}

func (r *InMemoryBookRepository) updatedBook(id int, updateBook Book) bool {
	for i, book := range books {
		if book.ID == id {
			books[i] = updateBook
			books[i].UpdatedAt = time.Now().Format(time.RFC3339)
			books[i].Finished = books[i].PageCount == books[i].ReadPage
			return true
		}
	}
	return false
}

func (r *InMemoryBookRepository) DeleteBook(id int) bool {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}

func (r *InMemoryBookRepository) FilterBooks(name, reading, finished string) []Book {
	filteredBooks := books

	if name != "" {
		filteredBooks = filteredBooksByName(filteredBooks, name)
	}
	if reading != "" {
		filteredBooks = filteredBooksByReading(filteredBooks, reading)
	}
	if finished != "" {
		filteredBooks = filteredBooksByFinished(filteredBooks, finished)
	}

	return filteredBooks
}

func filteredBooksByName(books []Book, name string) []Book {
	var result []Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(name)) {
			result = append(result, book)
		}
	}
	return result
}

func filteredBooksByReading(books []Book, reading string) []Book {
	var result []Book
	isReading := reading == "1"
	for _, book := range books {
		if book.Reading == isReading {
			result = append(result, book)
		}
	}
	return result
}

func filteredBooksByFinished(books []Book, finished string) []Book {
	var result []Book
	isFinished := finished == "1"
	for _, book := range books {
		if book.Reading == isFinished {
			result = append(result, book)
		}
	}
	return result
}
