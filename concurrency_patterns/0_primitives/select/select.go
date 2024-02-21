package main

import "fmt"

// select statement = lets a go routine wait on multiple communitcation operations

func main() {
	myChannel := make(chan string) // create channel
	anotherChannel := make(chan string) // create channel

	go func() {                    // fork child 1
		myChannel <- "data"
	}()

	go func() {                    // fork child 2
		anotherChannel <- "DATA"
	}()

	// select statement blocks until one case can run
	// if multiple messages are ready at the same time -> case block will be selected at random

	select {                       
	case msgFromMyChannel := <- myChannel:
		fmt.Println(msgFromMyChannel)

	case msgFromAnotherChannel := <- anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}