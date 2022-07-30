package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// function returns hello
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!!",
		"Great to see you, %v!",
		"HI, %v good to see you",
	}
	printSlice(formats)

	return formats[rand.Intn(len(formats))]
}

func printSlice(s []string) {
	fmt.Printf("len =%d  cap=%d %v \n", len(s), cap(s), s)

}
