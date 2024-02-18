package greetings

import (
	"fmt"
	"math/rand"
)

// Multiple greetings
var Formats = []string{
	"Nice to meet you %v",
	"What's up %v",
	"How's your family %v",
	"Hi, welcome there %v !",
	"Great to see you %v :)",
}

// This functions is to send greetins for one name
func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// This function return multiple greetings for more than two names
func MultipleGreetings(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// This functions send a random greeting
func randomFormat() string {

	return Formats[rand.Intn(len(Formats))]
}
