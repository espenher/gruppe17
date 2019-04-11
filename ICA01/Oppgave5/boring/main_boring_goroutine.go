package main

import (
	"fmt"

	"./boring"
)

// Kode fra Rob Pike https://talks.golang.org/2012/concurrency.slide#20
func main() {
	c := make(chan string)
	go boring.Boring10("boring!", c)
	for i := 0; ; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
	}
	fmt.Println("You're boring; I'm leaving.")
}
