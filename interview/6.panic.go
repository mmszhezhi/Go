package main

import (
	"fmt"
	"sync"
)

func main() {

	var w = sync.WaitGroup{}
	w.Add(2)
	//defer fmt.Println("fuck panic")
	go func() {
		defer w.Done()
		defer func() {
			recover()
			fmt.Println("fuck recover")

		}()
		panic("fuck")
	}()
	go func() {
		defer w.Done()
		fmt.Println("uuuu")
	}()
	w.Wait()
	fmt.Println("done ~~~~~~~~~~~~~~")
}
