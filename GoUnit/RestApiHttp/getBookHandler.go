package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Book struct {
	BookId     string `json:"bookId"`
	BookName   string `json:"bookName"`
	BookAuthor string `json:"bookAuthor"`
}

var books []Book

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/getBook/")
	for _, book := range books {
		if book.BookId == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found with this id ", http.StatusNotFound)
}

func main() {
	books = append(books, Book{"101", "Book Name 1", "Book Author 1"})
	books = append(books, Book{"102", "Book Name 2", "Book Author 2"})

	http.HandleFunc("/getBook", GetAllBook)
	http.HandleFunc("/getBook/", GetBookById)
	http.ListenAndServe(":8080", nil)
}
