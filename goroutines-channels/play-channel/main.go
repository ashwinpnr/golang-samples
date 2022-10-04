package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	c2 := make(chan int)
	go print_element2(c2, c)
	go print_element(c)

	//time.Sleep(5 * time.Second)
	x := <-c2
	fmt.Printf("value :%d \n", x)
}

func print_element(c chan int) {
	fmt.Println("in print_element")
	c <- 1
	close(c)

}
func print_element2(c2 chan int, c chan int) {
	<-c
	fmt.Println("in print_element2")
	c2 <- 2
	close(c2)
}
