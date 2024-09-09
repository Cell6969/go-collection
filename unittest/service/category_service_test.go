package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-unittest/entity"
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

func TestCategoryService_FindById2(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "aldo",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.FindById("2")
	assert.Nil(t, err)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
