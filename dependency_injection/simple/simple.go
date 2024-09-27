package simple

import "errors"

// create repository provider
type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

// create service provider
type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed to create service")
	} else {
		return &SimpleService{
			repository,
		}, nil
	}
}
