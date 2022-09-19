package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hellp(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
