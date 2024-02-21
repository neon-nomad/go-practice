package main

import "fmt"

func main() {
	myChannel := make(chan string) // create channel

	go func() {                    // fork child
		myChannel <- "data"
	}()

	msg := <-myChannel             // join point

	fmt.Println(msg)               
}