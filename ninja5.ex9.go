package main

import (
	"fmt"
)

func hmm(f func(string) string, s string) string {
	fmt.Printf("%T\t, %v\n", s, s)
	fmt.Println("inside hmm s = ", s)
	return f(s)
}
func main() {
	g := hmm(func(s string) string {
		return s
	}, "hello")
	fmt.Printf("%T\t, %v\n", g, g)
	fmt.Println("Inside main g = ", g)
	//	fmt.Println(hmm(say() string))
}
