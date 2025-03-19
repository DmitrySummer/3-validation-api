package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserEmail    string
	UserPassword string
	UserHost     string
	UserPort     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		UserEmail:    os.Getenv("EMAIL"),
		UserPassword: os.Getenv("PASSWORD"),
		UserHost:     os.Getenv("HOST"),
		UserPort:     os.Getenv("PORT"),
	}
}
