package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	fmt.Println("Hello World")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
