package greetings

import "fmt"

// Hello is a function that takes a string variable "name" that returns a string type
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name) // Sprintf is a string formatting function, where %v is placeholder
	return message
}
