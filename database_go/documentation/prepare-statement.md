# Prepare Statement

Di golang juga bisa melakukan preparation statement untuk query sql. Sebagai contoh:
```go
func TestPreparationStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "jo" + strconv.Itoa(i) + "@gmail.com"
		comment := "test comment" + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert new comment", id)
	}
}
```

Dengan demikian dengan query yang sama yang dibinding dengan pool yang sama maka bisa menjalankan perulangan.