package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	var err error
	var message string
	if name == "" {
		err = errors.New("Empty name!")
	} else {
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	}
	return message, err
}
