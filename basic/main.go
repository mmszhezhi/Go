package main

import (
	"Go/basic/utils"
	"fmt"
)

func main() {

	fmt.Println(utils.Add(utils.CONN_HOST, utils.CONN_PORT))
	fmt.Println(utils.Add(ConnHost, ConnPort))

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
