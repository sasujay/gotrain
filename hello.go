package main

import "fmt"

var x string = "Hello"

var s string
var i int
var f float32

func main() {
	n, _ := fmt.Println("Hello World", 42, true)
	fmt.Println(x)
	x = "Bye"
	fmt.Println(x)
	fmt.Println(n)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(s)
	foo()
	foo()
}
func foo() {
	fmt.Println(x)
}
