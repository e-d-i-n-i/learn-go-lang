package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel for strings
	messageChannel := make(chan string)

	// Launch a goroutine that sends a message into the channel
	go func() {
		// Simulate some work
		time.Sleep(1 * time.Second)
		messageChannel <- "Hello from goroutine!"
	}()

	// Receive the message from the channel (blocks until available)
	msg := <-messageChannel
	fmt.Println(msg)

	fmt.Println("Main function finished.")
}