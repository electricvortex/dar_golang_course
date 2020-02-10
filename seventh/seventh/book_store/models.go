package book_store


type BookStore interface {

	Create(book *Book) (*Book, error)

	GetBook(id int64) (*Book, error)

	ListBooks() ([]*Book, error)

	UpdateBook(id int64, book *Book) (*Book, error)

	DeleteBook(id int64)  error

	SaveBooks(filepath string) error
}


type Book struct {
	ID int64 `json:"id"`
	Title string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Description string `json:"description,omitempty"`
	NumberOfPages int `json:"number_of_pages,omitempty"`
}

type ConfigFile struct{
	JsonFilePath string `json:"filepath"`
	Port string `json:"port"`
}