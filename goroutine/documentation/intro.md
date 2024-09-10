# Goroutine
Goroutine bisa dibilang adalah versi mini dari thread atau mini thread yang berjalan secara asynchronus.
Cara membuat goroutine
```go
package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Done")

	time.Sleep(1 * time.Second)
}
```

Contoh lain untuk membuktikan apakah goroutine overload
```go
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
```