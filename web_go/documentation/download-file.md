# Download File

Selain upload file, golang juga mendukung untuk melakukan download file.

Contoh implementasi:
```go
package web_go

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	var file = r.URL.Query().Get("file")
	var download = r.URL.Query().Get("download")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	if download == "" {
		w.Header().Set("Content-Disposition", "inline;")
		http.ServeFile(w, r, "./resources/"+file)
	} else if download == "false" {
		http.ServeFile(w, r, "./resources/"+file)
	} else {
		w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
		http.ServeFile(w, r, "./resources/"+file)
	}
}

func TestDownloadFile(t *testing.T) {
	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
```