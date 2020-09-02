package main

import (
	"fmt"
)

func main() {

	//s := foo("sum", 2, 3, 4, 5, 6, 7, 8, 9)
	s := foo("hello")
	fmt.Println("Sum is", s)
}

func foo(su string, x ...int) int {
	fmt.Println(x, su)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Printf("%T\n", x)
	s := 0
	for _, v := range x {
		s = s + v
	}
	return s
}
