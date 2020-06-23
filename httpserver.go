package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":2343")
	defer listener.Close()
	con, _ := listener.Accept()
	defer con.Close()
	buf := make([]byte, 1024*4)
	n, _ := con.Read(buf)
	fmt.Println(string(buf[:n]))
}
