package services

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_DSN")

	if dsn == "" {
		log.Fatal("DATABASE_DSN environment variable is not set")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}
func GetDatabase() *gorm.DB {
	if DB == nil {
		ConnectDatabase()
	}
	return DB
}
