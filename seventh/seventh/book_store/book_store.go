package book_store

import (
	"bufio"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func NewBookStore(filename string) (BookStore, error) {
	file := &os.File{}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
	} else {
		file, err = os.Open(filename)
		if  err != nil {
			return nil, err
		}
	}

	buffer := bufio.NewReader(file)
	data, err := ioutil.ReadAll(buffer)
	if err != nil {
		return nil, err
	}
	books := []*Book{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &books); err != nil {
			return nil, err
		}
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
	file, err := os.OpenFile(filename, os.O_WRONLY, 777)
	if err != nil {
		return err
	}
	data, err := json.Marshal(bsc.books)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func (bsc *bookStoreClass) ListBooks () ([]*Book, error){
	return bsc.books, nil
}

func (bsc *bookStoreClass) UpdateBook(id string, book *Book) (*Book, error) {
	for _, v := range bsc.books {
		if v.ID == id {
			v = book
			return v, nil
		}
	}
	return nil, errors.New("Not found")
}

func (bsc *bookStoreClass) DeleteBook(id string) error{
	for i,v := range bsc.books {
		if v.ID == id {
			bsc.books = append(bsc.books[:i], bsc.books[i+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}