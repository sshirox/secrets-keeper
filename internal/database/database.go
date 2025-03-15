package database

import (
	"fmt"
	"log"
	"os"

	"github.com/sshirox/secrets-keeper/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.VaultSecret{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	DB = db
	fmt.Println("The database is connected")
}
