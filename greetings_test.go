package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) { // accepts pointer to a testing.T
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`) // panics when cannot compile
	msg, err := Hello(name)
	// where msg doesn't match expected `want` or an error is returned
	// a fatal fail is called with a message.
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want match "", error`, msg, err)
	}
}
