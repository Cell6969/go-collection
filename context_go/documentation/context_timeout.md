# Context Timeout
Singkatnya Context Cancel tapi secara otomatis karena ada waktu timeout

Contoh implementasinya:
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
				time.Sleep(1 * time.Second) // simulasi proses lambat
			}
		}
	}()

	return destination
}
```

Kemudian pada unit testnya:
```go
func TestContextWithTimeout(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("total goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n)
	}

	time.Sleep(1 * time.Second) // menunggu goroutine mati
	fmt.Println("total goroutine", runtime.NumGoroutine())
}
```

Jadi kiat set berapa lama timeoutnya kemudian defer close() untuk memastikan jika proses lebih cepat dari wktu timeout maka akan tetap diberhentikan