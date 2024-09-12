# SQl Parameter
Jikalau parameter langsung dihardcode maka akan terjadi nya sql injection.

Sebagai contoh, dibuat table user:
```sql
CREATE TABLE user(
    username VARCHAR(100) NOT NULL ,
    password varchar(100) NOT NULL
)ENGINE = InnoDB;

INSERT INTO user(username, password) values('admin', 'admin');
```

Kemudian pada code golang:
```go
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username from user WHERE username='" + username +
		"' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Gagal login")
	}
}
```

Jikalau username nya benar maka sukses login namun sebaliknya jika gagal maka gagal login. Namun jikalau dari user menginput parameter injeksi akan mengakibatkan terjadinya sql injection.

Bestnya adalah menggunakan parameter
## SQL Parameter
untuk implementasi parameter dari sql bisa seperti ini:
```go
func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password =? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Gagal login")
	}
}
```

Dengan demikian query injection tidak memungkinkan. Contoh lain untuk penggunaan parameter sql
```go

```