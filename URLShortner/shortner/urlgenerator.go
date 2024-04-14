package shortner

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

/*
This struct defines a URL shortener, which maintains a mapping between
shortened URLs and their original url
*/
type URLShortener struct {
	URL map[string]string
}

// method handles the logic for shortening URLs
func (us *URLShortener) HandleURLShorten(w http.ResponseWriter, r *http.Request) {

	// checks if the HTTP request method is POST.
	// If it's not, it returns a "405 Method Not Allowed" error.
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// retrieves the original URL from the form values of the HTTP request.
	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	/*
		 It iterates through the existing shortened URLs to check if the original URL
		 has already been shortened. If it finds a match, it renders an HTML response with
		the existing shortened URL and exits the function
	*/
	for shortKey, storedURL := range us.URL {
		if storedURL == originalURL {
			// Render the HTML response with the existing shortened URL
			shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)
			responseHTML := fmt.Sprintf(`
                <h2>URL Shortener</h2>
                <p>Original URL: %s</p>
                <p>Shortened URL: <a href="%s">%s</a></p>
                <form method="post" action="/shorten">
                    <input type="text" name="url" placeholder="Enter a URL">
                    <input type="submit" value="Shorten">
                </form>
            `, originalURL, shortenedURL, shortenedURL)
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, responseHTML)
			return
		}
	}

	/*
		generates a unique short key using the generateShortKey function
		and adds an entry to the URL map with the short key as the key and the
		original URL as the value.
	*/
	shortKey := generateShortKey()
	us.URL[shortKey] = originalURL

	// Construct the full shortened URL
	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	// Render the HTML response with the shortened URL
	responseHTML := fmt.Sprintf(`
        <h2>URL Shortener</h2>
        <p>Original URL: %s</p>
        <p>Shortened URL: <a href="%s">%s</a></p>
        <form method="post" action="/shorten">
            <input type="text" name="url" placeholder="Enter a URL">
            <input type="submit" value="Shorten">
        </form>
    `, originalURL, shortenedURL, shortenedURL)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, responseHTML)
}

/*
This function generates a random short key of length 6 characters using a
predefined character set consisting of lowercase and uppercase letters along with digits
*/
func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
