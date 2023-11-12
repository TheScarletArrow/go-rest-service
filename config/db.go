package config

import (
	"go-rest-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err2 := db.AutoMigrate(&models.User{})
	if err2 != nil {
		panic(err2)
	}
	DB = db
}
