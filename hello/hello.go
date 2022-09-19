package main

import (
	"fmt"

	"example/greetings"
)

func main() {
	message := greetings.Hello("NameA")
	fmt.Println(message)
}
