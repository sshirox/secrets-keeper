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
	Type     string `json:"type"`
	Data     string `json:"data"`
	Metadata string `json:"metadata"`
}

func AddVaultSecret(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	var req VaultRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		http.Error(w, "Data type id required", http.StatusBadRequest)
		return
	}

	encryptedData, err := auth.EncryptData(req.Data)
	if err != nil {
		http.Error(w, "Encryption error", http.StatusInternalServerError)
		return
	}

	secret := models.VaultSecret{
		UserID:        userID,
		Type:          req.Type,
		EncryptedData: []byte(encryptedData),
		Metadata:      req.Metadata,
	}

	database.DB.Create(&secret)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data saved"})
}

func GetVaultSecrets(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	var secrets []models.VaultSecret
	database.DB.Where("user_id = ?", userID).Find(&secrets)

	var decryptedSecrets []map[string]string
	for _, secret := range secrets {
		decryptedData, err := auth.DecryptData(string(secret.EncryptedData))
		if err != nil {
			continue
		}

		decryptedSecrets = append(decryptedSecrets, map[string]string{
			"id":       secret.ID,
			"type":     secret.Type,
			"data":     decryptedData,
			"metadata": secret.Metadata,
		})
	}

	json.NewEncoder(w).Encode(decryptedSecrets)
}

func DeleteVaultSecret(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	secretID := chi.URLParam(r, "id")

	result := database.DB.Where("id = ? AND user_id = ?", secretID, userID).Delete(&models.VaultSecret{})
	if result.RowsAffected == 0 {
		http.Error(w, "Secret not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Secret deleted"})
}
