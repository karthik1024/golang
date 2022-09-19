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

	message, err := greetings.Hello("Snoopy")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
