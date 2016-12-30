package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// listen
	fmt.Println("hello ant server")
	ln, err := net.Listen("tcp", ":3806")
	if err != nil {
		fmt.Println(err)
		panic("listen error")
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("recv error")
			conn.Close()
			return
		}
		str := string(data[0:n])
		fmt.Println(str)
	}
}
