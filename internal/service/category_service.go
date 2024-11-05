package service

import (
	"api-payment/internal/database"
	"api-payment/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {

	categories, err := cs.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {

	cat, err := cs.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (cs *CategoryService) PostCategories(cat *entity.Category) (string, error) {
	_, err := cs.CategoryDB.PostCategories(cat)
	if err != nil {
		return "", err
	}
	return cat.ID, nil

}
