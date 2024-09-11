# Context
Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout, dan sinyal deadline. Context biasanya dibuat per request (misal untuk server web melalui http request). Hampir semua bagian golang memanfaatkan context baik database, http server, http client, dan lain lain.

Contoh pembuatan context sederhana:
```go
package context_go

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
```

By default, ketika aplikasi nya berupa web maka secara otomatis akan dibuatkan context untuk http request.