package database

import (
	"api-payment/internal/entity"
	"database/sql"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cat *CategoryDB) GetCategories() ([]*entity.Category, error) {

	rows, err := cat.db.Query("select * from categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (c *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var cat entity.Category
	err := c.db.QueryRow("select * from categories where id = ?", id).Scan(&cat.ID, &cat.Name)
	if err != nil {
		return nil, err
	}
	return &cat, nil

}

func (cat *CategoryDB) PostCategories(category *entity.Category) (*entity.Category, error) {

	_, err := cat.db.Exec("INSERT INTO categories (id, name) VALUES (?,?)", category.ID, category.Name)
	if err != nil {
		return nil, err
	}
	return category, nil

}
