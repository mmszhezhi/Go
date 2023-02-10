package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	fmt.Println(&Student{Name: "menglu"} == &Student{Name: "menglu"})
	fmt.Println(Student{Name: "menglu"} == Student{Name: "menglu"})
	fuck()
}
func fuck() {
	fmt.Println([...]string{"1"} == [...]string{"1"})
	//fmt.Println([]string{"1"} == []string{"1"})
}
