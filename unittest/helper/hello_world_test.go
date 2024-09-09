package helper

import "testing"

func TestHelloWord(t *testing.T) {
	result := HelloWord("jonathan")
	if result != "Hello jonathan" {
		panic("Result it not match")
	}
}
