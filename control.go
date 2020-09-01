package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("wont print")
	default:
		fmt.Println("default print")
	case false:
		fmt.Println("print")
	}
}
