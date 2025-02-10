package main

import (
	"fmt"

	"LearningGo/greetings"
)

func main() {

	message := greetings.Hello("Gladys") // 'Hello' is a functioned wrote with a capital first letter -> acts like a public function in Java/ can be exported and used in other packages.

	fmt.Println(message)
}
