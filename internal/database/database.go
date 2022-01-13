package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database")

	//dbUsername := os.Getenv("DB_USERNAME")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbHost := os.Getenv("DB_HOST")
	//dbNAME := os.Getenv("DB_NAME")
	//dbPort := os.Getenv("DB_PORT")
	dsn := os.Getenv("DB_DSN")

	//dsn := fmt.Sprintf("host%s port=%s user=%s dbname%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbNAME, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open("postgres", dsn)

	if err != nil {
		return db, err
	}

	return nil, nil
}
