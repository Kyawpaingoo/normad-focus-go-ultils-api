package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	conStr := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(conStr), &gorm.Config{})

	if err != nil {
		log.Println("Error connecting to database")
	}

	DB = db
	log.Println("Connected to PostgreSQL")
}
