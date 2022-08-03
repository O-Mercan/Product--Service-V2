package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type DBConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBTable    string
	DBPort     string
}

func NewDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env vars")
	}
	return &DBConfig{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBTable:    os.Getenv("DB_TABLE"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
