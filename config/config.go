package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ApiAddr    string
	ApiPort    string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		ApiAddr:    os.Getenv("API_ADDR"),
		ApiPort:    os.Getenv("API_PORT"),
	}

	return cfg, nil
}
