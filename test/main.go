package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GreetingResponse struct {
	Greeting string `json:"greeting"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	var response GreetingResponse

	name := r.URL.Query().Get("name")
	if name != "" {
		response.Greeting = fmt.Sprintf("Hello, %s !", name)
	} else {
		response.Greeting = fmt.Sprintf("Hello User")
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/greet", greetHandler)

	log.Printf("greeter running on :8888")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Printf("shutting down: %v", err)
	}

}
