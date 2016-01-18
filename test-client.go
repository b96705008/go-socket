lpackage main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("connect to server error!")
		return
	}
	defer conn.Close()

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// send to socket
		_, _ = conn.Write([]byte(text))
		//fmt.Fprintf(conn, text + "\n")

		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if (message == "Good bye\n") {
			break
		} else {
			fmt.Print("Message from server: " + message)
		}
	}
	fmt.Println("conn is closed")
}
