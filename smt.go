package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating a buffered channel with a capacity of 3
	channel := make(chan int, 3)

	// Producer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			channel <- i // Send number to channel
			fmt.Println("Produced:", i)
		}
		close(channel) // Close the channel when done producing
	}()

	// Consumer goroutine
	go func() {
		for num := range channel {
			fmt.Println("Consumed:", num)
			time.Sleep(time.Second) // Simulate time-consuming task
		}
	}()

	// Wait for a moment before exiting
	time.Sleep(6 * time.Second)
}
