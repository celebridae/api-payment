package router

import (
	"api-payment/internal/database"
	"api-payment/internal/service"
	"api-payment/internal/webserver"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitializeRoute(db *sql.DB) *chi.Mux {

	categoryDB := database.NewCategoryDB(db)
	productDB := database.NewProductDB(db)

	categoryService := service.NewCategoryService(*categoryDB)
	productService := service.NewProductService(*productDB)

	produWebservice := webserver.NewProductWebserver(*productService)
	catWebservice := webserver.NewCategoryWebserver(*categoryService)

	// Router
	router := chi.NewRouter()
	// setter all log
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/products/{id}", produWebservice.GetProduct)
	router.Get("/products", produWebservice.GetProducties)
	router.Get("/products/category/{id}", produWebservice.GetProductByCategory)
	router.Post("/products", produWebservice.PostProduct)

	// category
	router.Get("/category", catWebservice.GetCategories)
	router.Get("/category/{id}", catWebservice.GetCategoryById)
	router.Post("/category", catWebservice.PostCategory)

	return router

}
