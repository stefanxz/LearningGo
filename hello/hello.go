package main

import (
	"LearningGo/greetings"
	"fmt"
	"log" //log is just a console.log that also prints the exact time at which the message is logged
)

func main() {

	log.SetPrefix("greetings: ") // instead of having a prefix as a date, it tells you where the error happened
	log.SetFlags(0)

	message, err := greetings.Hello("") // 'Hello' is a functioned wrote with a capital first letter -> acts like a public function in Java/ can be exported and used in other packages.
	//since greetings.Hello can now return an error value as well, it s return value is assigned to a (message, err) tuple
	if err != nil { // if the error variable gets a value, it means that an error was encountered -> err != nil
		log.Fatal(err) // fatal error, exists program -> everything below it is unreachable
	}

	//everything goes well, just print message.

	fmt.Println(message)
}

//I can create more loggers just for seeing what errors go on and where
//appLogger := log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
//dbLogger := log.New(os.Stdout, "DB: ", log.Ldate|log.Ltime|log.Lshortfile)
//netLogger := log.New(os.Stdout, "NET: ", log.Ldate|log.Ltime|log.Lshortfile)
