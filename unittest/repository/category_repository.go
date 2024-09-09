package repository

import "golang-unittest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
