package greetings

import (
	"errors"
	"fmt"
)

// Hello is a function that takes a string variable "name" that returns a string type
func Hello(name string) (string, error) { // now function can return error string type and error type
	if name == "" {
		return "", errors.New("this name was empty") // errors are never to be capitalized by convention
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name) // Sprintf is a string formatting function, where %v is placeholder
	return message, nil                              // return nil("") for when error is encountered
}
