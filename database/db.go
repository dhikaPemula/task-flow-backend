package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// "github.com/AndhikaBhas/miniProject.git/models"
)

var DB *gorm.DB

func Connect(databaseURL string) {
	var err error

	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	log.Println("Database connected")

	// err = DB.AutoMigrate(
	// 	&models.User{},
	// 	&models.Category{},
	// 	&models.Task{},
	// )
	
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
