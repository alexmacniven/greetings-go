package greetings

import "fmt"

// Function has a parameter called name of type string
// And a return type of string
func Hello(name string) string {
	// The := operator initiates and assigns a value to a variable
	// Sprintf returns a string with formatting applied
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// Return the message, which is the same type as the return signature
	return message
}
