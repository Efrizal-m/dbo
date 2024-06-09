package config

import (
	"dbo/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load environtment")
	} else {
		fmt.Println("succes load environtment")
	}

	dsn := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)
	fmt.Println(dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(&models.Customer{}, &models.Order{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
}
