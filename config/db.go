package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "postgresql://NomadDB_owner:npg_s46SgotmkZRO@ep-twilight-voice-a1kwjtq6-pooler.ap-southeast-1.aws.neon.tech/NomadDB?sslmode=require&channel_binding=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // Show only warnings/errors
	})

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	DB = db
	fmt.Println("Successfully connected to database")
}
