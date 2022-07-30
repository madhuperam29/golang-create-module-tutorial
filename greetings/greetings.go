package greetings

import (
	"errors"
	"fmt"
)

// function returns hello
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
