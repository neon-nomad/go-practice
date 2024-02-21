package main

import (
	"fmt"
	"time"
)

func example_2() {
	// infinite looping go routine
	go func() {
		for {
			select {
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()

	time.Sleep(time.Second * 10)
}