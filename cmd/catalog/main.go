package main

import (
	"api-payment/internal/database"
	"api-payment/internal/service"
	"api-payment/internal/webserver"
	"fmt"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("My API-PAYMENT INIT")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/api-payment_db")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := initializeRoute(db)

	fmt.Println("Server is running on port 85")
	http.ListenAndServe(":333", router)

}

func initializeRoute(db *sql.DB) *chi.Mux {

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
