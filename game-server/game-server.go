package main

import (
	"fmt"

	"github.com/mvangent/greetings"
)

func main() {
	message := greetings.Hello("Georg")
	fmt.Print(message)
}
