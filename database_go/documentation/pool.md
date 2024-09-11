# Pool

By default sql.Db di golang bukan koneksi ke database melainkan pooling atau database pooling.

Beberapa pengaturan database pooling di golang

| Method                       | Keterangan |
|------------------------------| -----------|
| (DB) SetMaxIdleConns(number) | Pengaturan berapa jumlah koneksi minimal yang dibuat |
| (DB) SetMaxOpenConns(number) | Pengaturan berapa jumlah koneksi maksimal yang dibuat |
| (DB) SetConnMaxIdleTime(duration) | Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus |
| (DB) SetConnMaxLifeTime(duration) | Pengaturan berapa lama koneksi boleh digunakan |


Contoh implementasi:
buat file database.go
```go
package database_go

import (
	"database/sql"
	"time"
)

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:33066)/test")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
```

Kode diatas menjadi landasan yang nantinya bakal dipakai disemua transaksi database.