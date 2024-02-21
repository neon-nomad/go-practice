package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

//entry point
func main() { // end

	// forked off main function; asynchronous  
	go someFunc("1") 
	go someFunc("2") 
	go someFunc("3")

	time.Sleep(time.Second * 2) // not representaive of actual joining process

	fmt.Println("hi")
} //start

// Fork-Join Model: see image
