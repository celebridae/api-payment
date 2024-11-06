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

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)
	//categoryDB.GetCategories()
	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	produWebservice := webserver.NewProductWebserver(*productService)
	catWebservice := webserver.NewCategoryWebserver(*categoryService)

	// Router
	c := chi.NewRouter()
	// setter all log
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/products/{id}", produWebservice.GetProduct)
	c.Get("/products", produWebservice.GetProducties)
	c.Get("/products/category/{id}", produWebservice.GetProductByCategory)
	c.Post("/products", produWebservice.PostProduct)

	// category
	c.Get("/category", catWebservice.GetCategories)
	c.Get("/category/{id}", catWebservice.GetCategoryById)
	c.Post("/category", catWebservice.PostCategory)

	fmt.Println("Server is running on port 85")

	http.ListenAndServe(":333", c)

}
