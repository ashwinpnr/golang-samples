package main

import (
	"fmt"
)

func check_function(c chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Checking couter : %d\n", i)
	}
	c <- true // return true to channel
}

func main() {
	channel1 := make(chan bool) // created channel
	fmt.Printf("Execution stage %d\n", 1)
	go check_function(channel1) // go routine
	fmt.Printf("Execution stage %d\n", 2)
	areWeDone := <-channel1 // got valur from channel
	fmt.Printf("check completed : %v\n", areWeDone)
}
