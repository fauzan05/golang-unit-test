package service

import (
	"fmt"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}} 
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_NotFound(t *testing.T) {
	// mock diprogram terlebih dahulu
	categoryRepository.Mock.On("FindById", 0).Return(nil)
	category, err := categoryService.Get(0)
	assert.Nil(t, category) // category bernilai nil
	assert.NotNil(t, err)   // eror bernilai error
	fmt.Println(err)
}

func TestCategoryService_Found(t *testing.T) {
	category := entity.Category{
		Id:   2,
		Name: "makanan",
	}
	categoryRepository.Mock.On("FindById", 2).Return(category) // membuat mock
	categoryRepository.Mock.On("HelloWorld").Return("Hello World")
	result, err := categoryService.Get(2) // mengeksekusi atau menjalankan method Get
	assert.Nil(t, err)
	assert.NotNil(t, result)
	fmt.Println(result)
}
