package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	// Create a list of all the IP adresses and ports you want to connect to
	addresses := []string{"127.0.0.1:5002", "127.0.0.1:5003", "127.0.0.1:5004", "127.0.0.1:5005", "127.0.0.1:5006"}

	for _, address := range addresses {
		go connectAndRequestUserInfo(address)
	}

	// Loop for ever until program shut down
	for {
	}
}

// Connect and start listening, then send user command
func connectAndRequestUserInfo(address string) {
	conn := connectTo(address)

	go listen(conn)

	send(conn, "user")
}

func connectTo(address string) net.Conn {
	// Connect to the given address via tcp
	conn, _ := net.Dial("tcp", address)

	// Return the connection
	return conn
}

func send(conn net.Conn, message string) {
	// Send string to server via the connection. Append \n as a message delimiter.
	fmt.Fprintf(conn, message+"\n")
}

func listen(conn net.Conn) {
	for {
		// Listen for text messages from the server, delimited by \n
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// If message isn't empty, print it.
		if message != "" {
			fmt.Printf(message)
		}
	}
}
