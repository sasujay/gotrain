package main

import (
	"fmt"
)

type square struct {
	len float64
}
type circle struct {
	rad float64
}

func (s square) area() float64 {
	a := s.len * s.len
	return a
}
func (c circle) area() float64 {
	a := 3.14 * c.rad * c.rad
	return a
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
func main() {
	s1 := square{
		len: 5,
	}
	c1 := circle{
		rad: 5,
	}
	info(s1)
	info(c1)
}
