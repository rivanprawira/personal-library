package models

import (
	"strings"
	"time"
)

// Membuat tipe data buku dengan struktur
type Book struct {
	ID         int
	Title      string
	Year       int
	Author     string
	PageCount  int
	ReadPage   int
	Finished   bool
	InsertedAt string
	UpdatedAt  string
}

// Membuat variabel books sebagai slice dari Book
// digunakan untuk menambahkan buku dalam Book(Perpusatakaan)
var books []Book

// Membuat variabel untuk ID buku
var lastID int = 0

// Interface untuk fitur dalam perpustakaan
type BookRepository interface {
	AddBook(book Book) Book
	GetAllBooks() []Book
	GetBookByID(id int) (*Book, bool)
	UpdateBook(id int, updatedBook Book) bool
	DeleteBook(id int) bool
	FilterBooks(name, finished string) []Book
}

// Membuat tipe data InMemoryBookRepository untuk implementasi interface sebelumnya
type InMemoryBookRepository struct{}

// Fungsi update buku.
func (r *InMemoryBookRepository) UpdateBook(id int, updatedBook Book) bool {
	for i, book := range books {
		if book.ID == id {
			// Hanya memperbarui field yang tidak kosong atau bernilai default
			if updatedBook.Title != "" {
				books[i].Title = updatedBook.Title
			}
			if updatedBook.Year != 0 {
				books[i].Year = updatedBook.Year
			}
			if updatedBook.Author != "" {
				books[i].Author = updatedBook.Author
			}
			if updatedBook.PageCount != 0 {
				books[i].PageCount = updatedBook.PageCount
			}
			if updatedBook.ReadPage != 0 {
				books[i].ReadPage = updatedBook.ReadPage
			}

			// Perbarui status buku
			books[i].Finished = books[i].PageCount == books[i].ReadPage
			books[i].UpdatedAt = time.Now().Format(time.RFC3339)

			return true
		}
	}
	return false
}

// Fungsi menambahkan Buku.
func (r *InMemoryBookRepository) AddBook(book Book) Book {
	lastID++
	book.ID = lastID
	book.InsertedAt = time.Now().Format(time.RFC3339)
	book.UpdatedAt = book.InsertedAt
	book.Finished = book.PageCount == book.ReadPage

	books = append(books, book)
	return book
}

// Fungsi mengambil semua buku.
func (r *InMemoryBookRepository) GetAllBooks() []Book {
	return books
}

// Fungsi Mengambil buku berdasarkan id.
func (r *InMemoryBookRepository) GetBookByID(id int) (*Book, bool) {
	for _, book := range books {
		if book.ID == id {
			return &book, true
		}
	}
	return nil, false
}

// Fungsi Menghapus Buku
func (r *InMemoryBookRepository) DeleteBook(id int) bool {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}

// Fungsi Filter Buku
func (r *InMemoryBookRepository) FilterBooks(name, finished string) []Book {
	filteredBooks := books

	if name != "" {
		filteredBooks = filteredBooksByName(filteredBooks, name)
	}
	if finished != "" {
		filteredBooks = filteredBooksByFinished(filteredBooks, finished)
	}

	return filteredBooks
}

// Fungsi filter buku dari nama
func filteredBooksByName(books []Book, name string) []Book {
	var result []Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(name)) {
			result = append(result, book)
		}
	}
	return result
}

// Fungsi filter buku dari finished
func filteredBooksByFinished(books []Book, finished string) []Book {
	var result []Book
	isFinished := finished == "1"
	for _, book := range books {
		if book.Finished == isFinished {
			result = append(result, book)
		}
	}
	return result
}
