package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sshirox/secrets-keeper/config"
	"net/http"
)

var client = resty.New().SetBaseURL("http://localhost:8081")

type Error struct {
	Message string `json:"message"`
}

func Register(email, password string) error {
	resp, err := client.R().
		SetBody(map[string]string{"email": email, "password": password}).
		Post("/register")

	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode() != http.StatusCreated {
		var apiErr Error
		if jsonErr := json.Unmarshal(resp.Body(), &apiErr); jsonErr == nil && apiErr.Message != "" {
			return fmt.Errorf("registration error: %s", apiErr.Message)
		}
		return fmt.Errorf("registration error: received status %d", resp.StatusCode())
	}

	return nil
}

func Login(email, password string) (string, error) {
	resp, err := client.R().
		SetBody(map[string]string{"email": email, "password": password}).
		Post("/login")

	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		var apiErr map[string]string
		if jsonErr := json.Unmarshal(resp.Body(), &apiErr); jsonErr == nil {
			if msg, exists := apiErr["message"]; exists {
				return "", fmt.Errorf("login error: %s", msg)
			}
		}
		return "", fmt.Errorf("login error: unexpected status code %d", resp.StatusCode())
	}

	var result map[string]string
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	token, exists := result["token"]
	if !exists {
		return "", errors.New("login error: missing token in response")
	}

	return token, nil
}

func GetVaultSecrets() ([]map[string]string, error) {
	token := config.GetToken()

	resp, err := client.R().
		SetAuthToken(token).
		Get("/vault")

	if err != nil {
		return nil, errors.New("error receiving data: " + err.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("server returned error: " + resp.String())
	}

	var secrets []map[string]string
	err = json.Unmarshal(resp.Body(), &secrets)
	if err != nil {
		return nil, errors.New("error parsing JSON response: " + err.Error())
	}

	return secrets, nil
}

func AddVaultSecret(secretType, data, metadata string) error {
	token := config.GetToken()

	resp, err := client.R().
		SetAuthToken(token).
		SetBody(map[string]string{"type": secretType, "data": data, "metadata": metadata}).
		Post("/vault")

	if err != nil || resp.StatusCode() != http.StatusCreated {
		return errors.New("addition data error: " + resp.String())
	}

	return nil
}
