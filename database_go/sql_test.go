package database_go

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

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

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "user'; DROP TABLE user; #"
	password := "user"

	query := "INSERT INTO user(username,password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}
