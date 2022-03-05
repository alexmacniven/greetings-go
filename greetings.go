package greetings

import (
	"errors"
	"fmt"
)

// Function has a parameter called name of type string
// And a return type of string and error
func Hello(name string) (string, error) {
	// When an empty value is supplied as name
	// Return an empty string and a new error
	if name == "" {
		return "", errors.New("empty name")
	}
	// The := operator initiates and assigns a value to a variable
	// Sprintf returns a string with formatting applied
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// Return the message and an empty error
	return message, nil
}
