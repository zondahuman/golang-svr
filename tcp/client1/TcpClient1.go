package main

import (
	"fmt"
	"net"
)

const (
	address = "localhost:2345"
)

func main() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("connect tcp server failure :", err.Error())
		return
	}
	fmt.Println("already connect tcp server")
	defer conn.Close()
	startClient(conn)
}

func startClient(conn net.Conn) {
	var message = "hello, who~!"
	for {
		conn.Write([]byte(message))
		buffer := make([]byte, 128)
		client, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("read tcp server exception", err.Error())
		}
		fmt.Println(string(buffer[0:client]))
	}
}
