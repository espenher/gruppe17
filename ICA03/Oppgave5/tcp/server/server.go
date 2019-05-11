package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

// Struct for holding name and email of a user.
type user struct {
	Name  string
	Email string
}

// Function for returning an instance of the user struct with the user's name and e-mail address.
func getUserJSON() string {
	user := &user{Name: "Hanna", Email: "someawesometestemail@gmail.com"}
	jsonBytes, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}

// Handler for handling new clients connecting
func handler(conn net.Conn) {

	// When first connecting, send "ok" back to the client.
	conn.Write([]byte("ok\n"))
	fmt.Println("A new client has connected!")

	// Loop for long-lived TCP connection.
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')

		if strings.HasPrefix(msg, "quit") {
			fmt.Println("Client quitting")
			break
		}

		if strings.HasPrefix(msg, "user") {
			fmt.Println("Client has requested user information.")
			result := getUserJSON() + "\n"
			conn.Write([]byte(string(result)))
		}
	}

	conn.Close()
}

func main() {
	createServer("5002")
}

func createServer(port string) {
	//Create a tcp listener
	listener, err := net.Listen("tcp", ":"+port)

	//Check for errors when starting the listener
	if err != nil {
		panic(err)
	}

	// Start accepting incoming connections. handle them when they do.
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handler(conn)
	}
}
