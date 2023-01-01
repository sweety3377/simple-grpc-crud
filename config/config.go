package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DbHost          string
	DbPort          string
	DbName          string
	DbUser          string
	DbPassword      string
	GrpcAddr        string
	GrpcPort        string
	GrpcGatewayAddr string
	GrpcGatewayPort string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		DbHost:          os.Getenv("DB_HOST"),
		DbPort:          os.Getenv("DB_PORT"),
		DbName:          os.Getenv("DB_NAME"),
		DbUser:          os.Getenv("DB_USER"),
		DbPassword:      os.Getenv("DB_PASSWORD"),
		GrpcAddr:        os.Getenv("GRPC_ADDR"),
		GrpcPort:        os.Getenv("GRPC_PORT"),
		GrpcGatewayAddr: os.Getenv("GRPC_GATEWAY_ADDR"),
		GrpcGatewayPort: os.Getenv("GRPC_GATEWAY_PORT"),
	}

	return cfg, nil
}
