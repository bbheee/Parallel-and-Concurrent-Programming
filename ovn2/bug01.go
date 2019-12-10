package main

import "fmt"

func main() {
	ch := make(chan string)
	//If there is no goroutine here, main will wait until something receives from the channel.
	//By adding a goroutine here, another thread will wait instead, while the main will recieve and print out from the channel.
	go func() {
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}
