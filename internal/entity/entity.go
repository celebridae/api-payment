package entity

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}

}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"image_url"`
}

func NewProduct(name string, description string, price float64, categoria string, image string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    categoria,
		ImageUrl:    image,
	}
}
