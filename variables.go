package main

import "fmt"

type sample int

var x sample
var y int
var b bool

const (
	p = 42
	q = 34.45
)
const (
	d1 = iota + 1
	d2
	d3
	d4
	d5
)
const (
	Jan = iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
	numberofMonths
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

const (
	l = 2016 + iota
	m = l + iota
	n = l + iota
)

func main() {
	fmt.Println(x)
	fmt.Println(l, m, n)
	fmt.Println(kb)
	fmt.Printf("%b\n", kb)
	fmt.Println(mb)
	fmt.Printf("%b\n", mb)
	fmt.Println(gb)
	fmt.Printf("%b\n", gb)
	x = 42
	s := `dkjfkdj
	ldkfldk
	 dfglkdf 
	 fgd;lgklf 
	 dlfgk`
	fmt.Println(s)
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Println(b)
	z := 42
	fmt.Println(y == z)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", q)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(d1, d2, d3, d4, d5)
	fmt.Println(Jan,
		Feb,
		Mar,
		Apr,
		May,
		Jun,
		Jul,
		Aug,
		Sep,
		Oct,
		Nov,
		Dec,
		numberofMonths)

}
