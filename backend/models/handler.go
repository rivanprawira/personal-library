package models

type BookHandler struct {
	Repo BookRepository
}

func NewBookHandler(repo BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}

func (h *BookHandler) AddBook(book Book) Book {
	return h.Repo.AddBook(book)
}

func (h *BookHandler) GetAllBooks() []Book {
	return h.Repo.GetAllBooks()
}

func (h *BookHandler) GetBookByID(id int) (*Book, bool) {
	return h.Repo.GetBookByID(id)
}

func (h *BookHandler) UpdateBook(id int, updatedBook Book) bool {
	return h.Repo.UpdateBook(id, updatedBook)
}

func (h *BookHandler) DeleteBook(id int) bool {
	return h.Repo.DeleteBook(id)
}

func (h *BookHandler) FilterBooks(name, finished string) []Book {
	return h.Repo.FilterBooks(name, finished)
}
