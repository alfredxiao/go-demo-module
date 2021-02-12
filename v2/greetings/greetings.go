package greetings

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("v2.2.2 %v, %v", name, quote.Go())
	return message, nil
}
