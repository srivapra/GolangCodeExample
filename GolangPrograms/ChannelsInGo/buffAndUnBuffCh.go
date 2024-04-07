/* UnBuffered Channel ->
1. Unbuffered channel have capacity 0.
2. When a sender sends a value to unbufferd channel. It blocks until reciever is ready
   to recieve the value from the channel.
3. Similarly, when a reciever attempts to recive a value
   from the channel it blocks until sender sends the value to the channel
*/

/* Buffered Channel ->
1. Buffered channel have capacity more than 0.
2. When a gorotine sends a value to bufferd channel. It doesn't blocks as long as the channel
   is not full. If the channel is full the sender blocks until space becomes available in the
   channel.
3. Similarly, when a goroutine attempts to recive the value from buffered channel. It doesn't block
   as long as channel is not empty. If the channel is empty the reciever blocks until value is
   sent to channel
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	ch1 := make(chan string, 2)

	/*
		This line is attempting to send a value to channel. However, since the channel is unbufferd
		and there is no other goroutine is ready to recive the value, this sends operation blocks
		and causes a deadlock
	*/
	//ch <- "Hi"

	// The deadlock will occur before meet this line
	//fmt.Println(<-ch)

	/*
		How we can fix this?
		The short answer is:
		you should either create a goroutine to send the data to the channel while having the main
		goroutine ready to receive, or you can use buffered channels to avoid blocking the sender
		until a receiver is ready. Below is the two approach we can use any one
	*/

	// In this version of the code, the sender is in a separate goroutine,
	// allowing the main goroutine to continue and receive the data from the channel.
	// Approach 1
	go func() {
		ch <- "hello"

	}()
	fmt.Println(<-ch)

	// Approach 2
	myChan := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Reciever is Ready before sending the value to channel
	go func() {
		defer wg.Done()
		msg := <-myChan
		fmt.Println(msg)

	}()

	// Sender sending the value to channel
	go func() {
		defer wg.Done()
		myChan <- 10
	}()
	wg.Wait()

	// ****************** Unbufferd Channel ***************************

	ch1 <- "My"   // Sender can send without blocking
	ch1 <- "Name" // Sender can send without blocking
	//ch <- "Is"    // Sender can block because chanel is full

	rmsg1 := <-ch1 // Reciver can recieve without blocking
	rmsg2 := <-ch1 // Reciver can recieve without blocking
	//rmsg3 := <-ch1 // Reciver can block because the channel is empty
	fmt.Println(rmsg1)
	fmt.Println(rmsg2)
	//fmt.Println(rmsg3)

}
