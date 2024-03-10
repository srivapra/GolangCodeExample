// Below is a basic example demonstrating how to create CRUD APIs using the http package in Go:

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Book represents a book entity.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	// Initialize some sample books
	books = append(books, Book{ID: "1", Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: "2", Title: "Book 2", Author: "Author 2"})

	// Define routes for CRUD operations
	http.HandleFunc("/books", getBooksHandler)
	http.HandleFunc("/books/create", createBookHandler)
	http.HandleFunc("/books/update/", updateBookHandler)
	http.HandleFunc("/books/delete/", deleteBookHandler)

	// Start HTTP server
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler function to get all books.
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Handler function to create a new book.
func createBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, newBook)
	w.WriteHeader(http.StatusCreated)
}

// Handler function to update a book.
func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract book ID from URL path
	id := strings.TrimPrefix(r.URL.Path, "/books/update/")
	var updatedBook Book
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.NotFound(w, r)
}

// Handler function to delete a book.
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract book ID from URL path
	id := strings.TrimPrefix(r.URL.Path, "/books/delete/")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.NotFound(w, r)
}
