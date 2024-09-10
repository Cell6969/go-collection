# Channel

Channel adalah tempat komunikasi secara synchronus yang bisa dilakukan oleh goroutine. Perlu diperhatikan saat mengirim data ke channel, goroutine akan terblock sampai ada yang menerima data tersebut. Jadi harus ada yg mengirim dan menerima dari channel.

Membuat channel:
```go
package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// masukkan data ke channel
	go func() {
		channel <- "lol"
	}()

	// mengambil data dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
```

## Channel Sebagai Parameter
Pada kasusnya, channel akan dilempar sebagai parameter di function dengan default pass by reference.

Contoh implementasinya:
```go
func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "forkey"
}

func TestChannelAsParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
```

## Channel In, Channel Out
Ada kondisi dimana ketika melempar channel sebagai parameter function , channel tersebut hanya untuk mengirim/menerima data saja. Hal ini bisa dilakukan dengan menandai channel in atau channel out.
Contoh implementasinya:
```go
func OnlySend(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "fore"
}

func OnlyReceive(channel <-chan string) {
	time.Sleep(1 * time.Second)
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlySend(channel)
	go OnlyReceive(channel)

	time.Sleep(5 * time.Second)
}
```

## Buffered Channel
Singkatnya buffered channel yaitu buffer yang digunakan untuk menampung data antrian di channel. Perlu diketahui di channel terdapat buffer capacity yang menentukan kapasitas data.

Contoh implementasinya:
```go
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // berapa data yang bisa dimasukkan
	defer close(channel)

	channel <- "leo"
	fmt.Println("Selesai")
}
```

Ketika menggunakan buffer, walaupun tidak ada yang mengambil data tetap akan berjalan berbeda dengan channel biasa ketika dijalankan akan menjadi deadlock.
Contoh menggunakan goroutine:
```go
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // berapa data yang bisa dimasukkan
	defer close(channel)

	go func() {
		channel <- "lo"
		channel <- "la"
		channel <- "be"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}
```
## Range Channel
Simpelnya ketika data yang diberikan tidak tau berapa jumlahnya biasanya dari iterasi maka dilakukan range channel.

Contoh implementasinya:
```go
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}
```
## Select Channel
Select channel berfungsi untuk memilih data tercepat dari beberapa channel.