package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	currency := 20
	count := 10
	fmt.Println("hello go")
	for i := 0; i < currency; i++ {
		go func() {
			for j := 0; j < count; j++ {
				sendMessage(i, j)
			}
		}()
	}
	select {}
}

func sendMessage(i, j int) {
	conn, err := net.Dial("tcp", "127.0.0.1:3806")
	if err != nil {
		panic("error")
	}
	message := "hello, i am thread :" + strconv.Itoa(i) + " dial:" + strconv.Itoa(j) + "\r\n"
	fmt.Fprint(conn, message)
}
