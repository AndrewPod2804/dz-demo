package main

import (
	"4-order-api/configs"
	"4-order-api/internal/auth"
	product2 "4-order-api/internal/product"
	"4-order-api/internal/user"
	"4-order-api/middleware"
	"4-order-api/pkg/db"
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

	phone := user.NewPhoneRepository(db)

	authService := auth.NewAuthService(phone)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	productRepository := product2.NewProductRepository(db)

	product2.NewProductHandler(router, product2.ProductHandlerDeps{
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
