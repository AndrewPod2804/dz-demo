package main

import (
	"4-order-api/internal/product"
	"4-order-api/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	//"os"
	//"os/user"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		//log.Fatal("Error loading .env file")
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&user.Phone{})
	if err != nil {
		panic(err)
	}

}
