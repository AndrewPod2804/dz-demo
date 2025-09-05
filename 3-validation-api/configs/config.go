package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Ver VerifyConfig
}

type VerifyConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		//log.Fatal("Error loading .env file")
		panic("Error loading .env file")
	}
	return &Config{
		Ver: VerifyConfig{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address:  os.Getenv("ADDRESS"),
		},
	}
}
