//粘包问题演示客户端
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn) {
	for i := 0; i < 1; i++ {
		words := "{\"Id\":1,\"UserName\":\"golang\",\"Message\":\" i go to school today\"}"
		conn.Write([]byte(words))
	}
}

func main() {
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("connect success")

	go sender(conn)
	buffer := make([]byte, 1024)
	result, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
		return
	}
	var output = string(buffer[:result])
	fmt.Println(conn.RemoteAddr().String(), "receive data string:", output)

	for {
		time.Sleep(1 * 1e9)
	}
}
