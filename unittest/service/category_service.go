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
