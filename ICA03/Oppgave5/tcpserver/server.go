package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

type user struct {
	Name  string
	Email string
}

func getUserJSON() string {
	user := &user{Name: "Hanna", Email: "someawesometestemail@gmail.com"}
	jsonBytes, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}

func handler(conn net.Conn) {
	conn.Write([]byte("ok\n"))

	// Loop for long-lived TCP connection.
	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')

		if strings.HasPrefix(msg, "quit") {
			fmt.Printf("Quitting")
			break
		}

		if strings.HasPrefix(msg, "user") {
			result := getUserJSON() + "\n"
			conn.Write([]byte(string(result)))
		}
	}

	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":5002")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handler(conn)
	}
}
