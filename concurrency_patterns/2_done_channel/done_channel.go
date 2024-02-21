package main

import (
	"fmt"
	"time"
)

// infinite looping go routine
func doWork(done <-chan bool) { // passing done channel as readonly 
	for {
		select {
		case <-done: // recieve close from main function
			return	
		default:
			fmt.Println("DOING WORK")
		}
	}
}

// main channel canceling a go routine
func main() { 
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)
	
	close(done)
}