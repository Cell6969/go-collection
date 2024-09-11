# GOMAXPROCS
GOMAXPROCS merupakan sebuah function dipackage runtime yang berfungsi untuk mengadjust jumlah thread.

Contoh penggunaan :
```go
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

```

Jika di running maka hasilnya seperti ini:
```shell
total cpue: 20
total thread: 20
total goroutine: 2
```

Jadi by default di komputer (saya) terdapat 20 CPU dengan 20 Thread. Namun ada 2 goroutine berjalan. 1 untuk merunning code program dan 1 untuk garbage collection.