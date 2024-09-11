# Context With Deadline
Mirip dengan timeout tapi langsung ditentukan kapan waktu cancelnya misal jam 12 siang, jam 1, seminggu dan lain lain.

Contoh implementasinya:
```go
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
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

