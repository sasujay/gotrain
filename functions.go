package main

import (
	"fmt"
)

func main() {

	//s := foo("sum", 2, 3, 4, 5, 6, 7, 8, 9)
	s := foo("hello")
	fmt.Println("Sum is", s)
	tot := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Sum is", tot)
	s9 := func() int { return 42 }
	fmt.Println("Anonymous Func:", s9())
	fmt.Printf("%T\t%v\n", s9, s9)
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
func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	tot := 0
	for _, v := range xi {
		tot += v
	}
	return tot
}
