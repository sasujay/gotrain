package main

import "fmt"

type person struct {
	fn  string
	ln  string
	age int
}

func changeme(p *person) {
	fmt.Println(p)
	fmt.Println(*p)
	/*p2 := person{
		fn:  "ad",
		ln:  "sj",
		age: 13,
	}*/
	//p = &p2
	fmt.Println(p)
	fmt.Println(*p)
	(*p).age = 14
	fmt.Println(*p)
}

func main() {
	p1 := person{
		fn:  "ab",
		ln:  "sj",
		age: 9,
	}
	fmt.Println("Main before", p1)
	fmt.Println(&p1)
	changeme(&p1)
	fmt.Println("Main after", p1)
	fmt.Println(&p1)
}
