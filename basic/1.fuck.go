package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

type Num interface {
	Integer | Float | string
}

func GMin[T constraints.Ordered](x, y T) T {

	return x + y
}

type Dictionay[K comparable, V any] map[K]V

type Student struct {
	Name  string
	Hobis []string
}

func main() {
	a := 4.5
	b := 5.6
	fmt.Println(GMin(a, b))
	fmt.Println(3 + 3.33)
	fmt.Println(GMin[string]("fuck", "u"))
	fmt.Println("b" > "aa")
	s1 := Student{
		Name: "fucker",
	}
	//s2 := Student{
	//	Name: "fuckerff",
	//}
	dict := Dictionay[Student, int]{s1: 1}
	//dict := Dictionay{s1: 1}
	fmt.Println(dict)
	//fmt.Println(s1>s2)
}
