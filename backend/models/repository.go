package models

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

// Interface untuk fitur dalam perpustakaan
type BookRepository interface {
	AddBook(book Book) Book
	GetAllBooks() []Book
	GetBookByID(id int) (*Book, bool)
	UpdateBook(id int, updatedBook Book) bool
	DeleteBook(id int) bool
	FilterBooks(name, finished string) []Book
}
