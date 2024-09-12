# Tipe Data

Tipe data untuk di database harus disesuaikan di golang . Berikut mappingan tipe data

| Tipe Data Database | Tipe Data Golang |
|--------------------|------------------|
| VARCHAR, CHAR      | string           |
| INT, BIGINT        | int32, int64     |
| FLOAT, DOUBLE      | float32, float64 |
| BOOLEAN            | bool             |
| DATE, DATETIME, TIME, TIMESTAMP | time.Time |              

Sebagai contoh, bisa table customer bisa diadjust sebagai berikut:
```sql
ALTER TABLE  customer
    ADD COLUMN email  VARCHAR(100),
    ADD COLUMN balance INT DEFAULT 0,
    ADD COLUMN rating DOUBLE DEFAULT 0.0,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married  BOOLEAN DEFAULT false;

DESC customer;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
values
    ('do', 'Do', 'do@gmail.com', 100000, 90.0,'1999-10-9', true),
    ('budi', 'Budi', 'budi@gmail.com', 150000, 85.5, '1986-06-10', true),
    ('jo', 'Jo', 'jo@gmail.com', 250000, 87.0, '1986-07-12', false);

SELECT * FROM customer;
```

Kemudian pada code golang:
```go
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at from customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id string
		var name string
		var email string
		var balance int32
		var rating float64
		var birth_date time.Time
		var created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("=======================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("email", email)
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		fmt.Println("birth_date", birth_date)
		fmt.Println("married", married)
		fmt.Println("created_at", created_at)
	}
}
```

Catatan: Untuk parse datetime perlu menambahkan pada url **parsedTime=true**

## Nullable Type
Ada kondisi dimana kolom data tersebut bisa jadi null, untuk menghandle tersebut harus menggunakan tipe data dari sql.

Updated:
```go
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at from customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id string
		var name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("=======================")
		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("email", email)
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		fmt.Println("birth_date", birth_date)
		fmt.Println("married", married)
		fmt.Println("created_at", created_at)
	}
}
```

Jikalau ada data yang null maka data tersebut akan menjadi false.