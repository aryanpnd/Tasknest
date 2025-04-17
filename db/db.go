package db

import (
	"fmt"
	"log"
	"os"

	"github.com/aryanpnd/Tasknest/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// dsn := "host=localhost user=postgres password=root dbname=tasknest_db port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to database")

	// Auto migrate Todo model
	err = DB.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	fmt.Println("✅ Database migrated")
}
