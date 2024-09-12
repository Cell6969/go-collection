# Auto Increment
Contoh implementasi auto increment untuk insert data di golang

Sebagai contoh , buat table comments
```sql
CREATE TABLE comments(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL ,
    comment TEXT,
    primary key (id)
)ENGINE =InnoDB;

DESC comments;
```

Kemudian pada golang:
```go
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "fauzi@gmail.com"
	comment := "Test Comment"

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment", insertId)
}
```

