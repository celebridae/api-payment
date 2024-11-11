package webserver

// webserver like controller

import (
	"api-payment/internal/entity"
	"api-payment/internal/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

type CategoryWebserver struct {
	CategoryService service.CategoryService
}

func NewCategoryWebserver(cs service.CategoryService) *CategoryWebserver {
	return &CategoryWebserver{CategoryService: cs}
}

// TODO: return all categories
func (c *CategoryWebserver) GetCategories(resp http.ResponseWriter, req *http.Request) {
	categories, err := c.CategoryService.GetCategories()
	if err != nil {
		//http.Error(resp, err.Error(), http.StatusInternalServerError)
		writeJSONResponse(resp, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	//json.NewEncoder(resp).Encode(categories)
	writeJSONResponse(resp, categories, http.StatusOK)
}

func (c *CategoryWebserver) GetCategoryById(resp http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		http.Error(resp, "id is required", http.StatusBadRequest)
		return
	}

	cat, err := c.CategoryService.GetCategory(id)
	if err != nil {
		//http.Error(resp, err.Error(), http.StatusInternalServerError)
		writeJSONResponse(resp, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	//json.NewEncoder(resp).Encode(cat)
	writeJSONResponse(resp, cat, http.StatusOK)
}

func (c *CategoryWebserver) PostCategory(resp http.ResponseWriter, req *http.Request) {
	var cat entity.Category

	err := json.NewDecoder(req.Body).Decode(&cat)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	newCat := entity.NewCategory(cat.Name)
	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(newCat)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		http.Error(resp, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
		return
	}

	result, err := c.CategoryService.PostCategories(newCat)
	if err != nil {
		//http.Error(resp, err.Error(), http.StatusInternalServerError)
		writeJSONResponse(resp, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	//json.NewEncoder(resp).Encode(result)
	writeJSONResponse(resp, result, http.StatusCreated)

}
