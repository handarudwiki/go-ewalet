package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Name:     os.Getenv("DATABASE_NAME"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
		},
		Email: Email{
			Host:     os.Getenv("EMAIL_HOST"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Port:     os.Getenv("EMAIL_PORT"),
			Email:    os.Getenv("EMAIL_NAME"),
		},
		Redis: Redis{
			Addres:   os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}
}
