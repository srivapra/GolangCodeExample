package main

import (
	"URLShortner/shortner"
	"fmt"
	"net/http"
)

func main() {

	shortener := &shortner.URLShortener{
		URL: make(map[string]string),
	}

	http.HandleFunc("/", shortner.HandleForm)
	http.HandleFunc("/shorten", shortener.HandleShorten)
	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
