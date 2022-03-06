package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// init function called when a package is imported.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v. Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

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
	message := fmt.Sprintf(randomFormat(), name)
	// Return the message and an empty error
	return message, nil
}

// Function takes a slice parameter.
// Returns a map of string keys and string value pairs.
func Hellos(names []string) (map[string]string, error) {
	// Maps need to be initialised with 'make'
	messages := make(map[string]string)
	// Iterate over a slice using 'range'.
	for _, name := range names {
		message, err := Hello(name)
		// Exit quickly if one names returns an error.
		if err != nil {
			return nil, err
		}
		// Assign a value to a key using index notation.
		messages[name] = message
	}
	return messages, nil
}
