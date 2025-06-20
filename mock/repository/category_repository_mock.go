package repository

import (
	"belajar-golang-dasar/mock/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock interface {
	Mock mock.Mock
}

func (repository *CateCategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil;
	}

	category := arguments.Get(0).(entity.Category)

	return &category
}

