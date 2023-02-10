package main

import (
	"fmt"
	"sync"
	"time"
)

// 要求手写代码
// 要求sync.WaitGroup支持timeout功能
// 如果timeout到了超时时间返回true
// 如果WaitGroup自然结束返回false
func main() {
	defer fmt.Println("fuck main")

	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			x := <-close
			fmt.Println(fmt.Sprintf("return %d iter %d", x, num))
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		//fmt.Println("timeout exit")
		close(c)

	}
	//time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {

	ch := make(chan bool, 1)

	go time.AfterFunc(timeout, func() {
		defer fmt.Println("fuck1")
		ch <- true
	})

	go func() {
		defer fmt.Println("fuck2")
		wg.Wait()
		ch <- false
	}()
	defer fmt.Println("waittime out exit")
	return <-ch
}
