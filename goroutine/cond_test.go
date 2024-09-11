package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	cond.L.Lock() // lock seperti biasa
	cond.Wait()   // melakukan wait
	fmt.Println("Done", value)
	cond.L.Unlock() // melakukan unlock
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(1 * time.Second)
	//		cond.Signal() // memberikan sinyal untuk unlock
	//	}
	//}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Broadcast()
		}
	}()

	group.Wait()
}
