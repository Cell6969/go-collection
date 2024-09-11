package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("total cpue:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine:", totalGoroutine)
}
