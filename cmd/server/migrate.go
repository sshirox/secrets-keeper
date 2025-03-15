package main

import (
	"fmt"
	"github.com/sshirox/secrets-keeper/internal/database"
	"github.com/sshirox/secrets-keeper/internal/models"
	"log"
)

func main() {
	database.ConnectDatabase()
	err := database.DB.AutoMigrate(&models.VaultSecret{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}
	fmt.Println("Migration completed successfully!")
}
