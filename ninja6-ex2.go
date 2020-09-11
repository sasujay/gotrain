package main

import (
	"fmt"
)

func main() {
	bar2()
	defer bar2()
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(s...)
	fmt.Printf("%T\t%v\n", sum, sum)
	fmt.Println("foo sum = ", sum)
	tot := bar(s)
	fmt.Printf("%T\t%v\n", tot, tot)
	fmt.Println("bar tot =", tot)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func bar(y []int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum
}
func bar2() {
	fmt.Println("Deffered!!!")
}
