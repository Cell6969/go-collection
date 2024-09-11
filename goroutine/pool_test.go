package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Default"
		},
	}
	group := sync.WaitGroup{}
	pool.Put("a")
	pool.Put("b")
	pool.Put("c")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}
