// Below is an example of a simple Go program that sends an HTTP request to a server
// and receives the response

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.javastudypoint.com/"

func main() {

	fmt.Println("Handling Web Request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Status)

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println(content)

}
