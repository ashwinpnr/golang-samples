package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	s := "teststring"
	f := 1.45
	DiscoverType(i)
	DiscoverType(s)
	DiscoverType(nil)
	DiscoverType(f)

}

func DiscoverType(t any) {
	switch v := t.(type) {
	case string:
		fmt.Printf("String Found : %s \n", v)
	case int:
		fmt.Printf("Integer Found : %d \n", v)
	default:
		var_type := reflect.TypeOf(t)
		if var_type == nil {
			fmt.Println("Type is nil")
			return
		}
		fmt.Printf("Dafault execution : %s \n", var_type)
	}

}
