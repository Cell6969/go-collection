# Context Cancel
singkatnya untuk memberikan sinyal cancel ke context.
Contoh kasusnya biasanya karena ada goroutine yang leak.

Contoh ada sebuah function :
```go
func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}
```

function tersebut bertujuan untuk menambahkan data secara incremental ke channel. Kemudian pada unit test:
```go
func TestContextWithCancel(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	destination := CreateCounter()

	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("total goroutine", runtime.NumGoroutine())
}
```

Jadi pada unit test , cek berapa goroutine yang berjalan diawal kemudian diakhir. Idealnya jumlah goroutine diawal dan diakhir akan selalu sama namun ketika di run hasilnya:
```shell
total goroutine 2
counter 1
counter 2
counter 3
counter 4
counter 5
counter 6
counter 7
counter 8
counter 9
counter 10
total goroutine 3
```

Ada 1 goroutine yang tetap berjalan walaupun tidak ada code program. Hal ini yang disebut sebagai goroutine leak. Hal ini terjadi karena pada goroutine function tidak berhenti untuk memasukkan data walapun pada main function sudah dibreak. Oleh karena itu perlu signal cancel context.

Contoh handlenya:
```go
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}
```
Pada function sisipkan context dan implement select case dimana untuk kondisi ctx yang done. Kemudian pada unit test nya:
```go
func TestContextWithCancel(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	cancel()
	time.Sleep(1 * time.Second) // menunggu goroutine mati
	fmt.Println("total goroutine", runtime.NumGoroutine())
}
```
Dengan code diatas kita mengirim sinyal cancel ke function ketika proses sudah selesai.