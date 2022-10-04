package main

import (
	"fmt"
	"sync"
)

func check_function(w *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Checking couter : %d\n", i)
	}
	w.Done()
}

func main() {
	var w sync.WaitGroup
	fmt.Printf("Execution stage %d\n", 1)
	w.Add(1)
	go check_function(&w) // go routine
	w.Wait()
	fmt.Printf("Execution stage %d\n", 2)
	//time.Sleep(20 * time.Second)

}
