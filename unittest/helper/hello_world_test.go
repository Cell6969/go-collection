package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTableHelloWorld(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "jonathan",
			request:  "jonathan",
			expected: "Hello jonathan",
		},
		{
			name:     "aldo",
			request:  "aldo",
			expected: "Hello aldo",
		},
		{
			name:     "tharin",
			request:  "tharin",
			expected: "Hello tharin",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWord(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("jonathan", func(t *testing.T) {
		result := HelloWord("Jonathan")
		assert.Equal(t, "Hello Jonathan", result)
	})
	t.Run("aldo", func(t *testing.T) {
		result := HelloWord("Aldo")
		assert.Equal(t, "Hello Aldo", result)
	})
}

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
