package shortner

import "net/http"

// Method to redirect short url to original url
func (us *URLShortener) HandleURLRedirect(w http.ResponseWriter, r *http.Request) {

	/*
		It extracts the shortened key from the request URL path. The path is expected
		to be in the format /short/{shortened_key}. It then trims the /short/ prefix to get
		the actual short key. If the short key is empty, it means that the request URL
		does not contain a valid shortened key, so it returns a "400 Bad Request" error.
	*/
	shortKey := r.URL.Path[len("/short/"):]
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	/*
	 retrieve the original URL corresponding to the shortened key
	 from the URL map maintained by the URLShortener instance. If the shortened key is not
	 found in the map, it means that the shortened URL does not exist in the system,
	 so it returns a "404 Not Found" error.
	*/
	originalURL, found := us.URL[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	// redirects the user to the original URL using the http.Redirect function.
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
