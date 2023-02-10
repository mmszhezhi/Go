package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

func min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

func comps[T constraints.Ordered](s []T) T {
	var result T
	for _, v := range s {
		result += v
	}
	return result
}

func main() {
	//a := 3
	//b := 4.3
	//fmt.Println(comps(a, b))

	fmt.Println(comps([]int{10, 2, 4, 1, 6, 8, 2}))
	fmt.Println(min([]int{10, 2, 4, 1, 6, 8, 2}))
	fmt.Println(max([]float64{3.2, 5.1, 6.2, 7.6, 8.2, 1.5, 4.8}))
}
