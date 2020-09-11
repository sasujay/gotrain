package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &b)
	fmt.Println(b)
	*b = 45
	fmt.Println(*b)
	fmt.Println(**&b)
	fmt.Println(*&b)
	fmt.Println(a)
	fmt.Println(&a)
	x := 77
	fmt.Println("Main before", x)
	fmt.Println("Main before", &x)
	foo(&x)
	fmt.Println("Main after", x)
	fmt.Println("Main after", &x)

}

func foo(y *int) {

	fmt.Println("In foo y= ", y)
	fmt.Println("in foo *y=", *y)
	*y = 88
	fmt.Println("In foo y= ", y)
	fmt.Println("in foo *y=", *y)

}
