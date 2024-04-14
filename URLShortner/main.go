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
	http.HandleFunc("/shorten", shortener.HandleURLShorten)
	http.HandleFunc("/short/", shortener.HandleURLRedirect)
	http.HandleFunc("/metrics", shortner.GetTopDomains)
	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
