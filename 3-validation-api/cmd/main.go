package main

import (
	"HomeVal3/configs"
	"HomeVal3/internal/verify"
	//"crypto/sha1"
	//"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("3-validation-api")
	config := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config: config,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("Server stopped")
}
