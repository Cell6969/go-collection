package test

import (
	"dependency_injection/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	var simpleService *simple.SimpleService = simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
