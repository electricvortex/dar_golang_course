package book_store

import (
	"bufio"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func NewBookStore(filename string) (BookStore, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	buffer := bufio.NewReader(file)
	data, err := ioutil.ReadAll(buffer)
	if err != nil {
		return nil, err
	}
	var books []*Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, err
	}
	defer file.Close()
	return &bookStoreClass{ books }, nil
}

type bookStoreClass struct {
	books []*Book
}

func (bsc *bookStoreClass) GetBook(id string) (*Book, error) {
	for _, v := range bsc.books {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("Not found ")
}

func (bsc *bookStoreClass) Create(book *Book) (*Book, error) {
	bsc.books = append(bsc.books, book)
	return book, nil
}

func (bsc *bookStoreClass) SaveBooks(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(bsc.books)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
