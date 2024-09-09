package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWordRequire(t *testing.T) {
	result := HelloWord("jonathan")
	require.Equal(t, "Hello jonathan", result, "Result must be 'Hello jonathan'")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWord("jonathan")
	assert.Equal(t, "Hello jonathan", result)
}

func TestHelloWord(t *testing.T) {
	result := HelloWord("jonathan")
	if result != "Hello jonathan" {
		panic("Result it not match")
	}
}
