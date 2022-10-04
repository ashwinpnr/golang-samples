package main

import "fmt"

func main() {
	fmt.Printf("Integer : %d\n", plusOne(2))
	fmt.Printf("Float : %v\n", plusOne(1.5))
}

func plusOne[V int | float32 | int64 | float64](t V) V {
	return t + 1
}
