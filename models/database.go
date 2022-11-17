package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failure in loading env file")
	}
	dbuser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, password, dbname)
	DB1, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// fmt.Println(dbuser, password, dbname)
	if err != nil {
		log.Fatal("Connection Error")
	} else {
		fmt.Println("Connected")
	}

	DB = DB1
	DB.AutoMigrate(&User{})
}
