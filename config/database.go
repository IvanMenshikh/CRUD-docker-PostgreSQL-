package config

import (
	"Golang-start/basic/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=ivanMENSHIKH150% dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключится к базе данных %v", err)
	}
	DB.AutoMigrate(&domain.Message{})
}
