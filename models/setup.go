package models

import (
	"log"
	"os"

	// "gorm.io/driver/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// dsn := "host=localhost user=postgres password=123 dbname=anama port=5432 sslmode=disable "
	// databse, err := gorm.Open(postgres.Open(dsn))

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in env")
	}
	databse, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	databse.AutoMigrate(&Source{})
	databse.AutoMigrate(&Product_category{})
	databse.AutoMigrate(&Product{})
	databse.AutoMigrate(&Customer{})
	databse.AutoMigrate(&Payment{})
	databse.AutoMigrate(&Transaction{})
	databse.AutoMigrate(&Balance{})

	DB = databse

}
