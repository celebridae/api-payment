package webserver

// webserver like controller

import (
	"api-payment/internal/entity"
	"api-payment/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
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

func (c *CategoryWebserver) GetCategoryById(resp http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		http.Error(resp, "id is required", http.StatusBadRequest)
		return
	}

	cat, err := c.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resp).Encode(cat)
}

func (c *CategoryWebserver) PostCategory(resp http.ResponseWriter, req *http.Request) {
	var cat entity.Category

	err := json.NewDecoder(req.Body).Decode(&cat)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.CategoryService.PostCategories(&cat)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(resp).Encode(result)

}
