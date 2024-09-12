# Execute Query
Query akan dijalankan di golang menggunakan function (DB) ExecContext(context, sql, params).

Sebagai contoh buat database, dan table customer:
```sql
CREATE TABLE customer(
    id varchar(100) not null ,
    name varchar(100) not null ,
    primary key (id)
)ENGINE = InnoDB;
```

Buat function GetConnection Sehingga bisa dipake diseluruh code program untuk eksekuis query:
```go
package database_go

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:33066)/go_db")
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

Kemudian coba insert data menggunakan golang:
```go
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES ('do', 'do')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}
```

## Query SQL
Untuk operasi SQL yang tidak membutuhkan hasil, bisa menggunakan perintah Exec, namun jika membutuhkan result bisa menggunakan function(DB) Query Context(context, sql, params).

Contoh implementasi:
```go
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name from customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
```

### Rows
Hasil query function jika tidak terjadi error adalah data structs sql.Rows. Dari rows tersebut kita bisa melakukan iterasi untuk mengambil data .
```go
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name from customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("name", name)
	}
}
```

