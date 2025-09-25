package main

import (
	"4-order-api/configs"
	"4-order-api/middleware"
	"4-order-api/pkg/db"
	"4-order-api/product"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("4-order-api")
	middleware.Init()

	conf := configs.LoadConfig()
	fmt.Println(conf)
	db := db.NewDb(conf)
	router := http.NewServeMux()
	productRepository := product.NewProductRepository(db)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}
	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error")
		return
	}
}
