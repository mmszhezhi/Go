package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	c := make(chan int, 100)
	for i := 0; i < 100; i++ {
		c <- i
	}
	var w = sync.WaitGroup{}
	w.Add(2)
	//ws := []string{"a", "b"}
	a, b := 0, 0
	go op(&c, "a", &w, &a)
	go op(&c, "b", &w, &b)

	w.Wait()
	fmt.Println(fmt.Sprintf("~~~~~~~result %d", a+b))
}

func op(x *chan int, name string, w *sync.WaitGroup, count *int) {
	defer w.Done()
	for {
		select {
		case d := <-*x:
			fmt.Println(fmt.Sprintf("%s get %d", name, d))
			*count++
		case <-time.After(time.Second * 2):
			fmt.Println(fmt.Sprintf("%s done!!!", name))
			return

		}
	}

}
