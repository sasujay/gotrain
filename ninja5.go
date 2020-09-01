package main

import (
	"fmt"
)

type person struct {
	fn     string
	ln     string
	flavor []string
}

func main() {

	p1 := person{
		fn:     "AD",
		ln:     "SJ",
		flavor: []string{"Vanila", "Choco"},
	}
	p2 := person{
		fn:     "AB",
		ln:     "SU",
		flavor: []string{"Chocolate", "Mint"},
	}
	m := map[string]person{
		"SJ": p1,
		"SU": p2,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	for i, v := range p2.flavor {
		fmt.Println(i, v)
	}
	fmt.Println(m)
	for _, v := range m {
		fmt.Println(v.fn)
		for i, f := range v.flavor {
			fmt.Println(i, f)
		}
		fmt.Println("-------------")
	}
}
