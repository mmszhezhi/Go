package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	conn, err := net.Dial("tcp", CONN_HOST+":"+"3333")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		fmt.Println("start epoch")
		ReadNWrite(conn)
		time.Sleep(time.Second * 3)
	}

	conn.Close()
}

func ReadNWrite(conn net.Conn) {
	message := "Test Request - by client\n"
	_, write_err := conn.Write([]byte(message))
	if write_err != nil {
		fmt.Println("failed:", write_err)
		return
	}
	conn.(*net.TCPConn).CloseWrite()

	buf, read_err := ioutil.ReadAll(conn)
	if read_err != nil {
		fmt.Println("failed:", read_err)
		return
	}
	fmt.Println(string(buf))
}
