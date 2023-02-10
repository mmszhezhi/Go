package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	//s := new(Show)
	s2 := Show{
		Param{"ff": 1},
	}
	s2.Param["RMB"] = 10000
	fmt.Println(s2)

}
