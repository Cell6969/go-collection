# Transactional

Di golang  bisa juga melakukan query transaction. Sebagai contoh:
```go
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {
		email := "jo" + strconv.Itoa(i) + "@gmail.com"
		comment := "test comment" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert new comment", id)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
```
untuk melakukan rollback tinggal ubah menjadi **tx.Rollback()**