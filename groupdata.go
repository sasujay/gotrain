package main

import (
	"fmt"
)

func main() {
	y := []int{1, 2, 3, 4, 5}
	fmt.Println(y)
	fmt.Println(y[:3])
	fmt.Println(y[2:3])
	y = append(y, 6, 7, 8, 9, 10)
	z := []int{11, 22, 33, 44}
	fmt.Println(y)
	fmt.Println(z)
	z = append(z, y...)
	fmt.Println(z)
	a := append(z[:3], z[9:]...)
	fmt.Println(z)
	fmt.Println(a)
	b := [][]int{y, z}
	fmt.Println(b)
	fmt.Println(b[1][2])
	for i, v := range y {
		fmt.Println(i, v)
	}
	for j := 0; j < 5; j++ {
		fmt.Println(y[j])
	}
	m := map[string]int{
		"r1": 1234,
		"r2": 6789,
		"r3": 9754,
	}
	fmt.Println(m)
	fmt.Println(m["r2"])
	g, good := m["r1"]
	fmt.Println(g)
	fmt.Printf("%T\t%T\n", g, good)
	fmt.Println(good)
	if v, ok := m["r2"]; ok {
		fmt.Println(v)

	} else {
		fmt.Println(ok)
	}
	m["r4"] = 5634
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, v := range y {
		fmt.Println(k, v)
	}
	delete(m, "r1")
	fmt.Println(m)
}
