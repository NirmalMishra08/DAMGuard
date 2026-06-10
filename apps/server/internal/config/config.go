package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT                string
	DB_URL              string
	CLICKHOUSE_PORT     string
	CLICKHOUSE_HOST     string
	CLICKHOUSE_DATABASE string
	CLICKHOUSE_USERNAME string
	CLICKHOUSE_PASSWORD string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("not able to load .env file")
		return nil, err
	}

	cfg := &Config{
		PORT:   os.Getenv("PORT"),
		DB_URL: os.Getenv("POSTGRES_CONN"),
		CLICKHOUSE_PORT: os.Getenv("CLICKHOUSE_PORT"),
		CLICKHOUSE_HOST: os.Getenv("CLICKHOUSE_HOST"),
		CLICKHOUSE_DATABASE: os.Getenv("CLICKHOUSE_DATABASE"),
		CLICKHOUSE_USERNAME: os.Getenv("CLICKHOUSE_USERNAME"),
		CLICKHOUSE_PASSWORD: os.Getenv("CLICKHOUSE_PASSWORD"),
	}

	return cfg, nil

}
