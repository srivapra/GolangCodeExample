package main

import (
	"URLShortner/shortner"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", shortner.HandleForm)
	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
