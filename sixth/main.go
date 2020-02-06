package main

import (
	"encoding/json"
	"fmt"
	"github.com/electricvortex/dar_golang_course/sixth/book_store"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func main() {
	bookStore, err := book_store.NewBookStore("books.json")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Methods("GET").Path("/{id}").HandlerFunc(GetBook(bookStore, "id"))
	router.Methods("POST").Path("/").HandlerFunc(CreateBook(bookStore))
	router.Methods("GET").Path("/exit").HandlerFunc(ExitWithSave(bookStore))
	fmt.Println("Server started")
	http.ListenAndServe("0.0.0.0:8080", router)
}

func GetBook(store book_store.BookStore, idParam string) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars[idParam]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Book ID not found "))
			return
		}
		book, err := store.GetBook(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("I'm sorry: " + err.Error()))
			return
		}
		data, err := json.Marshal(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	}
}

func CreateBook(store book_store.BookStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		book := &book_store.Book{}
		if err := json.Unmarshal(data, book); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		result, err := store.Create(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		response, err := json.Marshal(result)
		if err != nil {
			w.Write([]byte("Error: " + err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response)
		w.WriteHeader(http.StatusCreated)
	}
}

func ExitWithSave(book book_store.BookStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := book.SaveBooks("books.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}