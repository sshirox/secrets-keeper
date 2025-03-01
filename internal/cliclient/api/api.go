package api

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/sshirox/secrets-keeper/config"
	"net/http"
)

var client = resty.New().SetBaseURL("http://localhost:8080")

func Register(email, password string) error {
	resp, err := client.R().
		SetBody(map[string]string{"email": email, "password": password}).
		Post("/register")

	if err != nil || resp.StatusCode() != 201 {
		return errors.New("registration error")
	}

	return nil
}

func Login(email, password string) (string, error) {
	resp, err := client.R().
		SetBody(map[string]string{"email": email, "password": password}).
		Post("/login")

	if err != nil || resp.StatusCode() != http.StatusOK {
		return "", errors.New("login error")
	}

	token := resp.Result().(map[string]string)["token"]
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
