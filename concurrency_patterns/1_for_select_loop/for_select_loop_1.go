package main

import "fmt"

func example_1() {
	//charChannel := make(chan string)     unbuffered channel
	
	charChannel := make(chan string, 3) // buffered channel
	chars := []string{"a","b","c"}      // slice of chars

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)
    
	// after closing the channel we can loop over the data stored in the channel
	for result := range charChannel {  
		fmt.Println(result)
	}
}