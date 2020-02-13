package book_store


type BookStore interface {

	Create(book *Book) (*Book, error)

	GetBook(id string) (*Book, error)

	ListBooks() ([]*Book, error)

	UpdateBook(id string, book *Book) (*Book, error)

	DeleteBook(id string)  error

	SaveBooks(filepath string) error
}


type Book struct {
	ID string `json:"id"`
	Title string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Description string `json:"description,omitempty"`
	NumberOfPages int `json:"number_of_pages,omitempty"`
}

type ConfigFile struct{
	JsonFilePath string `json:"filepath"`
	Port string `json:"port"`
}