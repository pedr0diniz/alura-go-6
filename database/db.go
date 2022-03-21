package database

import (
	"log"

	"github.com/pedr0diniz/alura-go-5/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	connectionString := "host=localhost user=root password=root dbname=root port=5439 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Database connection error")
	}
	DB.AutoMigrate(&models.Student{})
}
