package test

import (
	"dependency_injection/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}
