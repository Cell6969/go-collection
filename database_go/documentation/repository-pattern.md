# Repository Pattern
Design pattern yang umum digunakan pada aplikasi golang adalah berbasis repository pattern. Jadi biasany komponen - komponen yang ada adalah:
1. Service => handle logic aplikasi
2. Entity => Sebagai entitas object data pada database
3. Repository => sebagai collection untuk logic database. 

Sebagai contoh untuk repository pattern yaitu data comments:

Buat entity untuk comment:
```go
package entity

type Comment struct {
	Id      int32
	Email   string
	Comment string
}
```

Kemudian buat repository interface
```go
package repository

import (
	"context"
	"database_go/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
```

Kemudian implement dari interface yang sudah dibuat:
```go
package repository

import (
	"context"
	"database/sql"
	"database_go/entity"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func (repository commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	//TODO implement me
	query := "INSERT INTO comments(email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Email)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	//TODO implement me
	query := "SELECT id,email,comment FROM comments WHERE id=? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	//TODO implement me
	query := "SELECT id,email,comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
```
Kemudian implememtasi New Repository
```go
func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}
....
```

Dengan demikian repository sudah terimplement.

Kemudian untuk unit test nya:
```go

```