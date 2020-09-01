package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type student struct {
	person
	rollnum int
}

func main() {
	p1 := person{
		first: "AD",
		last:  "SJ",
		age:   13,
	}
	p2 := person{
		first: "AB",
		last:  "SJ",
		age:   9,
	}
	fmt.Println(p1)
	fmt.Printf("%v\t%v\n", p1.first, p1.age)
	fmt.Printf("%v\t%v\n", p2.first, p2.age)
	s1 := student{
		person: person{
			first: "AB",
			last:  "SJ",
			age:   9,
		},
		rollnum: 23,
	}

	fmt.Println(s1)
	fmt.Printf("%v\t%v\t%v\n", s1.first, s1.age, s1.rollnum)

}
