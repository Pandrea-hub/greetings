package greetings

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	name := "David"
	message, err := Hello(name)
	assert.Empty(t, err)
	assert.True(t, strings.Contains(message, name))
	assert.True(t, ContainsFormat(name))
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	assert.Empty(t, message)
	assert.EqualErrorf(t, err, "Empty name", "The error doesn't coincidence")
}

func ContainsFormat(name string) bool {
	message, _ := Hello(name)
	flag := false
	for _, n := range Formats {
		messageExpected := fmt.Sprintf(n, name)
		if messageExpected == message {
			flag = true
		}
	}
	return flag
}
