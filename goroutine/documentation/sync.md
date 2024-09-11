# Sync
Pada golang terdapat sebuah struct yaitu sync yang bertujuan untuk menghandle masalah race condition

## sync.Mutex
Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking maka tidak ada yg bisa melakukan locking lagi sampai di unlock.
Contoh penggunaan mutex:
```go
package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0

	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("x =", x)
}
```

## sync.RWMutex
Jika mutex hanya melock write data, maka RWMutext dapat dilakukan lock untuk read dan write. Contoh implementasinya:
```go
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // write
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock() // write
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Latest balance = ", account.GetBalance())
}
```
## sync.WaitGroup
WaitGroup adalah fitur yang digunakan untuk menunggu sebuah proses goroutine selesai.
Ada beberapa hal dalam menggunakan WaitGroup
1. Add(int) untuk menandai adanya proses goroutine.
2. Done() untuk memberi tahu bahwa proses goroutine selesai
3. Wait() untuk menunggu semua proses selesai.

Contoh implementasi:
```go
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
```
Jadi done ditaru di function yang dijalankan ole goroutine sedangkan add ditaro sebelum function dijalankan.

## sync.Once
fitur golang yang bisa digunakan untuk memastikan bahwa sebuah function dieksekusi hanya sekali. Artinya ketika ada banyak go routine maka hanya satu goroutine yang mengakses function tersebut sisanya dihiraukan.
Syarat bisa dilakukan once adalah function tersebut tidak memiliki parameter
Contoh implementasinya:
```go
package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
```
## sync.Pool
Pool adalah implementasi design pattern bernama object pool pattern. Singkatnya design pattern pool ini digunakan untuk menyimpan data selanjutnya untuk menggunakan datanya, bisa dengan mengambil dari pool dan setelah selesai menggunakan data tersebut bisa dikembalikan ke poolnya. Implementasi pool di golang sudah aman dari race condition.

Contoh implementasinya:
```go
package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}
	group := sync.WaitGroup{}
	pool.Put("a")
	pool.Put("b")
	pool.Put("c")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}
```

jikalau ditambahkan timeSleep pada function maka sebagian terjadi <nil>. Untuk memberi default value bisa dilakukan seperti ini

````go
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
````
## sync.Map
Sama seperti map biasa tapi aman untuk dilakukan concurrent menggunakan goroutine.

Contoh implementasinya:
```go
package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
```

## sync.Cond
Singkatnya adalah implementasi locking berbasis kondisi 

Contoh implementasinya:
```go
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

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // memberikan sinyal untuk unlock
		}
	}()

	group.Wait()
}
```

Dalam code diatas, Signal ditujukan untuk memberi tahu goroutine satu persatu. Untuk memberitahu semua goroutine sekaligus bisa menggunakan broadcast.
```go
go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Broadcast()
		}
	}()
```

## sync.Atomic
atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent. Note: ini untuk kasus data primitive. Contoh implementasinya:
```go
package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i <= 100_000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("x = ", x)
}

```