package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Connect() error
	GetDB() *gorm.DB
}

type PostgresDatabase struct {
	dsn             string
	timeoutConexion int
	db              *gorm.DB
}

func NewPostgresDatabase(dsn string, timeout int) *PostgresDatabase {
	return &PostgresDatabase{
		dsn:             dsn,
		timeoutConexion: timeout,
	}
}

func (p *PostgresDatabase) Connect() error {
	if p.dsn == "" {
		return fmt.Errorf("DATABASE_DSN environment variable is not set")
	}
	var conectada bool
	for !conectada {
		database, err := gorm.Open(postgres.Open(p.dsn), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database:", err)
			p.timeoutConexion--
			if p.timeoutConexion <= 0 {
				return fmt.Errorf("timeout reached. Could not connect to database")
			}
			time.Sleep(1 * time.Second)
			continue
		}
		conectada = true
		p.db = database
	}
	return nil
}

func (p *PostgresDatabase) GetDB() *gorm.DB {
	if p.db == nil {
		err := p.Connect()
		if err != nil {
			log.Fatal(err)
		}
	}
	return p.db
}

var DB Database

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_DSN")
	DB = NewPostgresDatabase(dsn, 60)
	err := DB.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase() *gorm.DB {
	return DB.GetDB()
}
