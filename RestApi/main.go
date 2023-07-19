package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string
	Name   string
	Author string
}

type Books []Book

var books = Books{
	{
		Id:     "101",
		Name:   "Java",
		Author: "abc",
	},
	{
		Id:     "102",
		Name:   "Python",
		Author: "abc",
	},
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error during creating the book records")
	}
	var book Book
	json.Unmarshal(reqBody, &book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func getAllBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)

}

func getBookById(w http.ResponseWriter, r *http.Request) {
	bookid := mux.Vars(r)["id"]

	for _, v := range books {
		if v.Id == bookid {
			json.NewEncoder(w).Encode(v)

		}
	}

}

func deleteBookById(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]

	for k, v := range books {
		if v.Id == bookId {
			books = append(books[:k], books[k+1:]...)
		}
	}
}

func updateBookById(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]
	var bookData Book
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error during updation")
	}
	json.Unmarshal(req, &bookData)

	for k, v := range books {
		if v.Id == bookId {
			v.Name = bookData.Name
			books = append(books[:k], v)
		}
		json.NewEncoder(w).Encode(bookData)
	}

}

func HandleRequest() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/createbook", CreateBook)
	myRouter.HandleFunc("/allBook", getAllBook)
	myRouter.HandleFunc("/book/{id}", getBookById)
	myRouter.HandleFunc("/deletebook/{id}", deleteBookById)
	myRouter.HandleFunc("/updatebook/{id}", updateBookById)
	log.Fatal(http.ListenAndServe(":9091", myRouter))
}

func main() {
	HandleRequest()
}
