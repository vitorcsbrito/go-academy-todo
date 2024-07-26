package main

import (
	"fmt"
	"time"
)

// Protected Data resource - not concurrent safe
var coffeeOrder = []byte("               ")

func main1() {

	go writeText("Choca          ")
	go writeText("Mocha with Milk")
	go writeText("Banana Shake   ")
	go writeText("Tea no sugar   ")

	// What on earth will be in coffeeOrder now?
	time.Sleep(1 * time.Second)

	fmt.Println(string(coffeeOrder))
}

func writeText(newOrder string) {

	// CRITICAL SECTION STARTS
	orderAsBytes := []byte(newOrder)
	for index, b := range orderAsBytes {
		coffeeOrder[index] = b
		time.Sleep(10 * time.Millisecond)
	}
	// CRITICAL SECTION ENDS
}
