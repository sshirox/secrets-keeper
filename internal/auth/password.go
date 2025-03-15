package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%s", salt, hashBase64), nil
}

func VerifyPassword(password, hash string) bool {
	parts := strings.Split(hash, "$")
	if len(parts) != 2 {
		return false
	}

	salt := parts[0]
	expectedHash := argon2.IDKey([]byte(password), []byte(salt), 1, 64*1024, 4, 32)
	expectedHashBase64 := base64.StdEncoding.EncodeToString(expectedHash)

	return expectedHashBase64 == parts[1]
}
