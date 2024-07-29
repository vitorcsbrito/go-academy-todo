package main

import (
	"fmt"
	"time"
)

// Protected Data resource - not concurrent safe
var coffeeOrder = []byte("               ")

func main1() {

	ch := make(chan []byte)

	go writeText("Choca          ", ch)
	_ = <-ch

	go writeText("Mocha with Milk", ch)
	_ = <-ch

	go writeText("Banana Shake   ", ch)
	_ = <-ch

	go writeText("Tea no sugar   ", ch)
	_ = <-ch

	// What on earth will be in coffeeOrder now?
	time.Sleep(1 * time.Second)

	fmt.Println(string(coffeeOrder))
}

func writeText(newOrder string, ch chan []byte) {

	// CRITICAL SECTION STARTS
	orderAsBytes := []byte(newOrder)
	for index, b := range orderAsBytes {
		coffeeOrder[index] = b
		time.Sleep(10 * time.Millisecond)
	}

	// CRITICAL SECTION ENDS
	ch <- orderAsBytes
}
