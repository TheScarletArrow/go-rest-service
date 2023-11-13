package config

import (
	"fmt"
	"go-rest-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	fmt.Println("Hostname:", hostname)
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@db-1:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err2 := db.AutoMigrate(&models.User{})
	if err2 != nil {
		panic(err2)
	}
	DB = db
}
