package main

import "fmt"

const (
	a = iota
	b = 4
	xx
	tt
)
const (
	e    = iota
	name = "menglu"
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(xx)
	fmt.Println(tt)
}
