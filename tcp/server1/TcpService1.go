package main

import (
	"fmt"
	"net"
)

const (
	ip   = "localhost"
	port = 2345
)

func main() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("已初始化连接，等待客户端连接.........")
	startServer(listen)
}

func startServer(listen *net.TCPListener) {
	for {

		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error)
			continue
		}
		fmt.Println("client connection from :", conn.RemoteAddr().String())
		defer conn.Close()
		go func() {
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data)
				fmt.Println("client send data :", string(data[0:i]))
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
			}
		}()

	}

}
