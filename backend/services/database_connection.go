package services

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var timeout_conexion = 60
var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_DSN")

	if dsn == "" {
		log.Fatal("DATABASE_DSN environment variable is not set")
	}
	var conectada bool
	for !conectada {
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database:", err)
			timeout_conexion--
			if timeout_conexion <= 0 {
				log.Fatal("Timeout reached. Could not connect to database.")
			}
			time.Sleep(1 * time.Second)
			continue
		}
		conectada = true
		DB = database
	}

}
func GetDatabase() *gorm.DB {
	if DB == nil {
		ConnectDatabase()
	}
	return DB
}
