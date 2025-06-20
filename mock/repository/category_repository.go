package repository

import "belajar-golang-dasar/mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}