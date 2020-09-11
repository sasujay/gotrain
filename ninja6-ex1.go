package main

import (
	"fmt"
)

func main() {
	x := foo()
	y, s := bar()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(s)
}

func foo() int {
	return 42
}
func bar() (int, string) {
	return 43, "hello"
}
