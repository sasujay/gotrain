package main

import (
	"fmt"
)

func main() {

	s := say()
	fmt.Printf("%T\t, %v\n", s, s)
	fmt.Println(s())

}

func say() func() string {
	return func() string {
		return "hello"
	}

}
