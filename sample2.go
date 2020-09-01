package main

import "fmt"

var y = 42
var a = 4.1
var z = "hello"

type swapna float32

var b swapna = 43.2453465

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", z)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	a = float64(b)
	fmt.Printf("%#T\n", a)
	fmt.Println(a)

}
