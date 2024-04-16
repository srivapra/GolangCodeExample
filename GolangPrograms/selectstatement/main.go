/*
In Go, the select statement is used for concurrent programming with channels. The select statement
is used to choose from multiple send/receive channel operations.
The select statement blocks until one of the send/receive operations is ready.
If multiple operations are ready, one of them is chosen at random. The syntax is similar to
switch except that each of the case statements will be a channel operation.
*/

/*
Let's say you are building a web service that needs to fetch data from multiple external APIs
to respond to a client request. Instead of making these requests sequentially, which
could significantly increase the response time, you can make them concurrently using goroutines
and channels, and then use select to wait for the responses. By utilizing select along with
goroutines and channels, the web server can fetch data from multiple APIs concurrently and
respond to client requests efficiently, reducing overall response time.
*/

package main

import (
	"fmt"
	"time"
)

func fetchDataFromAPI(apiURL string, result chan<- string) {
	// Simulate fetching data from an API (could be an HTTP request)
	time.Sleep(time.Second) // Simulate network latency

	// Assume data is fetched successfully
	result <- fmt.Sprintf("Data from %s", apiURL)
}

func main() {
	api1 := make(chan string)
	api2 := make(chan string)

	// Start fetching data from API 1 and API 2 concurrently
	go fetchDataFromAPI("https://api1.com/data", api1)
	go fetchDataFromAPI("https://api2.com/data", api2)

	// Use select to wait for data from either API
	select {
	case data := <-api1:
		fmt.Println(data)
	case data := <-api2:
		fmt.Println(data)
	}

	// Here you can continue processing the fetched data...
}
