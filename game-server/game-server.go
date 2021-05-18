package main

import (
	"fmt"
	"log"

	"github.com/mvangent/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Melvin")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(message)
}
