package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sshirox/secrets-keeper/internal/auth"
	"github.com/sshirox/secrets-keeper/internal/database"
	"github.com/sshirox/secrets-keeper/internal/models"
)

type VaultRequest struct {
	Data     string `json:"data"`
	Metadata string `json:"metadata"`
}

func AddVaultEntry(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	var req VaultRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	encryptedData, err := auth.EncryptData(req.Data)
	if err != nil {
		http.Error(w, "Encryption error", http.StatusInternalServerError)
		return
	}

	entry := models.VaultSecret{
		UserID:        userID,
		EncryptedData: []byte(encryptedData),
		Metadata:      req.Metadata,
	}

	database.DB.Create(&entry)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data saved"})
}

func GetVaultEntries(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	var entries []models.VaultSecret
	database.DB.Where("user_id = ?", userID).Find(&entries)

	decryptedEntries := []map[string]string{}
	for _, entry := range entries {
		decryptedData, err := auth.DecryptData(string(entry.EncryptedData))
		if err != nil {
			continue
		}

		decryptedEntries = append(decryptedEntries, map[string]string{
			"id":       entry.ID,
			"data":     decryptedData,
			"metadata": entry.Metadata,
		})
	}

	json.NewEncoder(w).Encode(decryptedEntries)
}

func DeleteVaultEntry(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	entryID := chi.URLParam(r, "id")

	result := database.DB.Where("id = ? AND user_id = ?", entryID, userID).Delete(&models.VaultSecret{})
	if result.RowsAffected == 0 {
		http.Error(w, "Secret not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Secret deleted"})
}
