package main

import (
	"net"
	"fmt"
)

func handleConn(conn net.Conn) {
	conn.Write([]byte("Welcome WPP server!" + "\n"))
	defer conn.Close()

	for {
		// read message
		buf := make([]byte, 1024);
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			break
		}

		// reply message
		message := string(buf[:n])
		fmt.Println("Message Received:", message)
		conn.Write(buf[:n])
	}
	fmt.Println("socket disconnect.")
}

func main() {
	fmt.Println("Launching WPP test server on 8081 port...")

	//listener
	ln, _ := net.Listen("tcp", ":8081")

	//process every connection
	for {
		conn, err := ln.Accept()
		if err != nil {
			conn.Close()
			continue
		}
		fmt.Println("Here comes new socket!")
		go handleConn(conn)
	}
}
