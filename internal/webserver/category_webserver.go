package webserver

// webserver like controller

import (
	"api-payment/internal/service"
	"encoding/json"
	"net/http"
)

type CategoryWebserver struct {
	CategoryService service.CategoryService
}

func (c *CategoryWebserver) NewCategoryWebserver(cs service.CategoryService) *CategoryWebserver {
	return &CategoryWebserver{CategoryService: cs}
}

// TODO: return all categories
func (c *CategoryWebserver) GetCategories(resp http.ResponseWriter, req *http.Request) {
	categories, err := c.CategoryService.GetCategories()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resp).Encode(categories)
}
