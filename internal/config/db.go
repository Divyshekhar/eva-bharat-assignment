package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	godotenv.Load()
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")
	DB_TIMEZONE := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
		DB_SSLMODE,
		DB_TIMEZONE,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database" + err.Error())
		return
	}
	log.Print("Database Connected Successfully")

}
