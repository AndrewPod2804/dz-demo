package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Bb DbConfig
	//Auth AuthConfig
}
type DbConfig struct {
	Dsn string
}

//type AuthConfig struct {
//	Secret string
//}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Bb: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		//Auth: AuthConfig{
		//	Secret: os.Getenv("SECRET"),
		//},
	}
}
