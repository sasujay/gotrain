package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Fn  string
	Ln  string
	Age int
}

func main() {

	p1 := person{
		Fn:  "AD",
		Ln:  "SJ",
		Age: 13,
	}
	p2 := person{
		Fn:  "AB",
		Ln:  "SJ",
		Age: 9,
	}
	fmt.Println(p1, p2)

	people := []person{p1, p2}
	fmt.Println(people)
	fmt.Printf("%T\n", people)
	b, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", err)
	fmt.Println(string(b))
	err := json.Marshal(b)

}
