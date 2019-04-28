package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Set up two endpoints.

	// One enpoint for a sort of "hello world"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to our user API!")
	})

	// One enpoint for delivering user info json
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, getUserJSON())
	})

	// Run as goroutine
	go http.ListenAndServe(":5010", nil)

	// Live forever
	for {
	}
}

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
