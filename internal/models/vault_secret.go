package models

import (
	"time"

	"gorm.io/gorm"
)

type VaultSecret struct {
	ID            string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID        string         `gorm:"not null"`
	Type          string         `gorm:"not null"`
	EncryptedData []byte         `gorm:"not null"`
	Metadata      string         `gorm:"type:jsonb"`
	CreatedAt     time.Time      `gorm:"default:now()"`
	UpdatedAt     time.Time      `gorm:"default:now()"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
