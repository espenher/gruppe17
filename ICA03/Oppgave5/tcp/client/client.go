package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Create connection to the server
	conn := connectTo("127.0.0.1:5002")

	for {
		// Listen for messages from the server
		go listen(conn)

		// Read text written in the console
		text := readConsole()

		// Send text written in the console to the server via the connection
		send(conn, text)
	}
}

func connectTo(address string) net.Conn {
	// Connect to the given address via tcp
	conn, _ := net.Dial("tcp", address)

	// Return the connection
	return conn
}

func readConsole() string {
	// Read input from Stdin
	reader := bufio.NewReader(os.Stdin)

	// Read as text delimited by newline
	text, _ := reader.ReadString('\n')

	// Return the text
	return text
}

func send(conn net.Conn, message string) {
	// Send string to server via the connection. Append \n as a message delimiter.
	fmt.Fprintf(conn, message+"\n")
}

func listen(conn net.Conn) {
	// Listen for text messages from the server, delimited by \n
	message, _ := bufio.NewReader(conn).ReadString('\n')

	// If message isn't empty, print it.
	if message != "" {
		fmt.Println(message)
	}
}
