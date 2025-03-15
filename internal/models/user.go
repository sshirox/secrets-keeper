package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string         `gorm:"unique;not null"`
	PasswordHash string         `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"default:now()"`
	UpdatedAt    time.Time      `gorm:"default:now()"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
