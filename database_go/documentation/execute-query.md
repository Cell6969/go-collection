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