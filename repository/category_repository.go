package repository

import "golang-unit-test/entity"

type CategoryRepository interface{
	FindById(id int) *entity.Category
	HelloWorld() string
}