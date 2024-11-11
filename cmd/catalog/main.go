package main

import (
	"api-payment/internal/router"
	"fmt"
	"net/http"

	"database/sql"

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

	router := router.InitializeRoute(db)

	fmt.Println("Server is running on port 333")
	http.ListenAndServe(":333", router)

}
