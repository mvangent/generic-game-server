package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name was given to greet")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
