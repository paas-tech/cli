package auth

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/paastech-cloud/cli/pkg/utils"
)

const base_url = "http://localhost:3000/auth"

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"accessToken"`
}

// Send a login request to the PaasTech API and returns the Access Token if successful
func Login(email string, password string) (string, error) {
	// Create JSON request body
	request, err := json.Marshal(loginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	// POST request to API
	resp, err := http.Post(base_url+"/login", "application/json", bytes.NewReader(request))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response is an error
	err = utils.Error(resp)
	if err != nil {
		return "", err
	}

	// Parse JSON body
	var jwt loginResponse
	err = json.NewDecoder(resp.Body).Decode(&jwt)
	if err != nil {
		return "", err
	}

	return jwt.AccessToken, err
}
