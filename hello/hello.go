package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	// Set Logger properties.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"A", "B", "C"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
