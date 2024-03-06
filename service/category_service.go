package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type CategoryService struct{
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id int) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category Not Found")
	} else {
		return category, nil
	}
}

func (helloworld CategoryService) GetHello() string{
	return "Hello World"
}