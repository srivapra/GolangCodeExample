/*
Context is a package in the standard library which is mainly used to propagate cancellation signals
from one function to another or even from micro service to another.

Let’s consider the example of a user sending a get request to a web server to download 50 images,
zip it and send the zipped response back. The request has been triggered from the browser and let’s
say it takes 70 seconds to complete. The user is not patient enough and decides to cancel it when it
is still being processed by the server. Wouldn’t it be nice if the server gets to know that the
user(client) has cancelled the request so that the server can also terminate the request and save
valuable CPU and memory? This is the perfect use case for a context. The context allows  the server
to know when a request has been cancelled by the client so that it can terminate its resources and move on.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

// func longRunning() int {
// 	count := 0
// 	for i := 0; i < 5; i++ {
// 		count = count + i
// 		fmt.Println("Current value of count:", count)
// 		time.Sleep(2 * time.Second)

// 	}
// 	return count
// }

// func main() {
// 	count := longRunning()
// 	fmt.Println("count is ", count)
// }

/*
In the above program, longRunning function runs a for loop for 5 iterations which sleeps for 2
seconds between each iteration. The code print output after 10 seconds. What if we want to terminate
the long running function after 5 seconds? This is possible with the help of context.
The main function will send a cancellation signal to longRunning function after 5 seconds
using the context.
*/

func longRunning(ctx context.Context) (int, error) {
	count := 0
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			count = count + i
			fmt.Println("Current value of count:", count)
			time.Sleep(2 * time.Second)
		}
	}
	return count, nil
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		time.Sleep(8 * time.Second)
		cancelFunc()
	}()
	count, err := longRunning(ctx)
	if err != nil {
		fmt.Println("long running task exited with error", err)
		return
	}

	fmt.Println("count is ", count)
}

/*
The above program will print the output and will terminate after 2 seconds.
This is exactly what we wanted. We have sent a termination signal to a long running program
after 2 seconds and have successfully stopped the long running function.
*/
