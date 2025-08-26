package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //DB pointer to gorm


func Connect() {
	dsn := fmt.Sprintf( // dsn meaning data source name
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // &grom.Config{} is default
	if err != nil {                                          // nill meaning null
		log.Fatal("Fail to connect database:", err)
	}

	DB = db
	fmt.Println("Connect to PostgreSQL")
}
