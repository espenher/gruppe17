package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn := connectTo("127.0.0.1:5000")

	for {
		go listen(conn)
		text := readConsole()
		send(conn, text)
	}
}

// func main() {
// 	addresses := []string{"127.0.0.1:5000", "127.0.0.1:5001"} //, "127.0.0.1:5002"} //, "127.0.0.1:5003"}

// 	for _, address := range addresses {
// 		go requestTheThing(address)
// 	}

// 	for {
// 		readConsole()
// 	}
// }

// func requestTheThing(address string) {
// 	conn := connectTo(address)

// 	send(conn, "user")
// 	listen(conn)
// }

func connectTo(address string) net.Conn {
	// connect to this socket
	conn, _ := net.Dial("tcp", address)
	return conn
}

func readConsole() string {
	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func send(conn net.Conn, message string) {
	// send to socket
	fmt.Fprintf(conn, message+"\n")
}

func listen(conn net.Conn) {
	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	if message != "" {
		fmt.Printf(message)
	}
}
