package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age)
}

func main() {
	p1 := person{
		first: "AD",
		last:  "SU",
		age:   13,
	}
	p1.speak()
}
