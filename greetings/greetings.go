package greetings

import (
	"fmt"

	"rsc.io/quote"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("v1.1.3 %v, %v", name, quote.Go())
	return message
}
