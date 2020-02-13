package book_store

import "testing"

var bookStore BookStore

func TestBookStoreClass_GetBook(t *testing.T) {
	var err error
	bookStore, err = NewBookStore("../books.json")
	if err != nil {
		t.Error(err)
	}
	book, err := bookStore.GetBook("10")
	if err != nil {
		t.Error(err)
	}
	if book.Title != "ASddsdasdad" {
		t.Error("Title is not equal")
	}
}

func TestBookStoreClass_Create(t *testing.T) {

}
