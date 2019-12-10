package main

import "fmt"
import "time"

func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	//time.Sleep() function waits for goroutine to be done before main is done.
	//But this will not work in all situations:
	//For example, if it takes a long time to process the data from the channel,
	//waiting for 10 milliseconds will not be enough.
	time.Sleep(time.Millisecond * 10)
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
