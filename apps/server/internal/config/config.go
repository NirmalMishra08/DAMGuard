package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT   string
	DB_URL string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("not able to load .env file")
		return nil, err
	}

	cfg:= &Config{
		PORT: os.Getenv("PORT"),
		DB_URL: os.Getenv("PORT"),
	}

	return cfg, nil


}
