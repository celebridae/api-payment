package database

import (
	"api-payment/internal/entity"
	"database/sql"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (db *ProductDB) GetProducties() ([]*entity.Product, error) {
	rows, err := db.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var prod entity.Product
		if err := rows.Scan(&prod.ID, &prod.Name, prod.Description, &prod.Price, prod.ImageUrl, prod.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &prod)
	}
	return products, nil
}

func (prod *ProductDB) GetProduct(id string) (*entity.Product, error) {

	var product entity.Product

	err := prod.db.QueryRow("select * from products where id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageUrl, &product.CategoryID)
	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (db *ProductDB) PostProduct(prod *entity.Product) (*entity.Product, error) {

	_, err := db.db.Exec("insert into products (id, name,description,price,category_id, image_url) VALUES (?,?,?,?,?,?)", prod.ID, prod.Name, prod.Description, prod.Price, prod.ImageUrl, prod.CategoryID)
	if err != nil {
		return nil, err
	}
	return prod, nil

}

func (prod *ProductDB) GetProductByCategory(idCategory string) ([]*entity.Product, error) {
	//var product entity.Product
	rows, err := prod.db.Query("select * from Products where category_id = ?", idCategory)
	//.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageUrl, &product.Category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var prod entity.Product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.ImageUrl, &prod.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &prod)
	}
	return products, nil

}
