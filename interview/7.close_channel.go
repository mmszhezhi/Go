package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	var w = sync.WaitGroup{}
	w.Add(1)

	go func() {

		defer w.Done()
		defer fmt.Println("over~~~~~~~~~~")
		fmt.Println("block")
		x := <-c
		fmt.Println("fuck c")
		fmt.Println(x)
	}()

	time.AfterFunc(time.Second*5, func() {
		fmt.Println("over time")
		close(c)
	})

	w.Wait()

}
