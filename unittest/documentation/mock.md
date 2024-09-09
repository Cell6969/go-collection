# Mock

Mock adalah object yang sudah dibuat dari awal untuk menyesuaikan ekspektasi.
Untuk melakukan mock alangkah lebih baiknya membentuk struct terlebih dahulu. Pada kasus implementasi ini akan melakukan query ke database.

1. Buat entity Category
```go
package entity

type Category struct {
	Id   string
	Name string
}
```
2. Buat interface repository
```go
package repository

import "golang-unittest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
```
3. Buat Service
```go
package service

import (
	"errors"
	"golang-unittest/entity"
	"golang-unittest/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) FindById(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}
```
4. Buat mock repository
```go
package repository

import (
	"github.com/stretchr/testify/mock"
	"golang-unittest/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
```
5. Buat test Category Service
```go
package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unittest/repository"
	"testing"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: &categoryRepository}

func TestCategoryService_FindById(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.FindById("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
```
Penjelasan:
Pertama kita buat interface untuk repository category, kemudian kita buat repository category sesuai interface. Selanjutnya kita buat service untuk category. Untuk testing repository category dikarenakan tidak connect ke db maka dibuat lah mock repository. Dari mock repository kita buat seolah
olah terdapat data. Kemudian dilakukan testing terhadap service category. Karena di service membutuhkan repository maka kita masukkan repository mock yang sebelumnya sudah dibuat. Selanjutnya tinggal mengatur logic pada unit test tersebut.

Conton test success:
```go
func TestCategoryService_FindById2(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "aldo",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.FindById("2")
	assert.Nil(t, err)
	assert.Equal(t, category, result)
}
```