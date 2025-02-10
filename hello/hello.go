package main

import (
	"fmt"

	"LearningGo/greetings"
)

func main() {

	message := greetings.Hello("Gladys")

	fmt.Println(message)
}
