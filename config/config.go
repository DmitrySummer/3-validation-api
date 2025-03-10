package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserEmail    string
	UserPassword string
	UserAddress  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error load .env file, use default config")
	}
	return &Config{
		UserEmail: os.Getenv("EMAIL"),
		UserPassword: os.Getenv("PASSWORD"),
		UserAddress: os.Getenv("ADDRESS"),
		},
}
