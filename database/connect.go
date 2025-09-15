package database

import (
	"os"
	"fmt"
	"log"
	"strconv"

	"github.com/waxer59/basic-go-fiber-api/config"
	"github.com/waxer59/basic-go-fiber-api/internal/upload/uploadModels"
	"github.com/waxer59/basic-go-fiber-api/internal/user/userModels"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	p := os.Getenv("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection established")

	DB.AutoMigrate(&userModels.User{}, &uploadModels.Upload{})
	fmt.Println("Database Migrated")
}
