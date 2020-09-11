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

func (s student) hello() {
	fmt.Println("I am", s.first, s.last)
}

func (a student) alive() {
	fmt.Println("I", a.first, a.last, "am a student")
}
func (b person) alive() {
	fmt.Println("I", b.first, b.last, "am a person")
}

type living interface {
	alive()
}

func bar(l living) {
	fmt.Println("bar called living", l)
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
	fmt.Println("*************")
	s1.hello()
	s1.alive()
	p1.alive()
	bar(s1)
	bar(p1)
	bar(p2)
}
