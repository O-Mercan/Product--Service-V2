package database

import (
	"fmt"
	"log"
	"os"

	//"github.com/O-Mercan/Product--Service-V2/util"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("cannot load env file:", err)
	}

	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUserName, dbPassword, dbTable, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	postgresDB, err := db.DB()
	if err := postgresDB.Ping(); err != nil {
		return db, err
	}
	return db, nil
}
