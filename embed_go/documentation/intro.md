# Golang Embed

Di golang terdapat package embed yaitu fitur baru untuk mempermudah membaca isi file pada saat compile time secara otomatis, biasanya digunakan untuk render html.


## Embed string

Untuk mengembed file menjadi variable string,  bisa dilakukan seperti ini

buat file misal txt
```text
1.0.0-SNAPSHOT
```

Kemudian pada golang:
```go
package embed_go

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string // harus diluar

func TestString(t *testing.T) {
	fmt.Println(version)
}
```

Catatan: Embed to String cocok untuk file file yang terbaca by string misal txt, csv, html

## Embed Byte

Selain embed ke string, bisa juga embed ke byte. Biasanya cocok untuk file yang berformat png,jpg, binary,dll

Contoh siapkan file gambar, kemudian pada code golang:
```go
//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
```

## Embed Multiple Files
Selain mengembed 1 file, bisa juga mengembed multiple file.

Sebagai contoh, buat multiple file dalam satu folder kemudian pada code golang:
```go
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
//go:embed files/d.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

	d, _ := files.ReadFile("files/d.txt")
	fmt.Println(string(d))
}
```

## Path Matcher
Sebelumnya file yang diembed merupakan exact match artinya sama persis, digolang bisa juga mengembed file dengan regex menggunaka path matcher.

Sebagai contoh:
```go
// membaca semua file yang memiliki extension .txt
//
//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
```

## Compile
File yang diembed dengan golang ketika di build akan dicompile menjadi binary. File tersebut sudah tidak bisa lagi diubah jikalau sudah dicompile. Untuk mengubah nya perlu dicompile lagi.

Sebagai contoh, pada main function:
```go

```