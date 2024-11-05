package webserver

import (
	"api-payment/internal/entity"
	"api-payment/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type ProductWebserver struct {
	ProductWebserver service.ProductService
}

func NewProductWebserver(ps service.ProductService) *ProductWebserver {
	return &ProductWebserver{ProductWebserver: ps}
}

func (pw *ProductWebserver) GetProducties(resp http.ResponseWriter, req *http.Request) {

	prod, err := pw.ProductWebserver.GetProducties()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resp).Encode(prod)

}

func (pw *ProductWebserver) GetProductByCategory(resp http.ResponseWriter, req *http.Request) {
	var prod entity.Product

	err := json.NewDecoder(req.Body).Decode(&prod)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := pw.ProductWebserver.GetProduct(prod.ID)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resp).Encode(result)
}

func (pw *ProductWebserver) PostProduct(resp http.ResponseWriter, req *http.Request) {

	var product entity.Product

	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := pw.ProductWebserver.PostProduct(product)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resp).Encode(result)
}

func (pw *ProductWebserver) GetProduct(resp http.ResponseWriter, req *http.Request) {
	var product entity.Product

	id := chi.URLParam(req, "id")
	if id != "" {
		http.Error(resp, "id is requirid", http.StatusBadRequest)
		return
	}

	result, err := pw.ProductWebserver.GetProduct(product.ID)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resp).Encode(result)

}
