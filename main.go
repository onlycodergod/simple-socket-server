package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")

	ln, _ := net.Listen("tcp", ":8000")

	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
	}
}
