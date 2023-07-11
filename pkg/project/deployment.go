package project

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/paastech-cloud/cli/pkg/utils"
)

type projectDeployRequest struct {
	EnvVars map[string]string `json:"env_vars"`
}

// Deploy a project to PaaSTech
func (p *Project) Deploy(baseURL string, accessToken string, envVars map[string]string) error {
	// Create JSON request body
	request, err := json.Marshal(projectDeployRequest{
		EnvVars: envVars,
	})
	if err != nil {
		return err
	}

	// Create PATCH request to API
	req, err := http.NewRequest(http.MethodPatch, baseURL+"/projects/"+p.Id+"/deploy", bytes.NewReader(request))
	if err != nil {
		return err
	}

	// Add access token as Bearer token in headers
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response is an error
	err = utils.Error(resp)
	if err != nil {
		return err
	}

	return nil
}

// Stop a project from PaaSTech
func (p *Project) Down(baseURL string, accessToken string) error {
	// Create POST request to API
	req, err := http.NewRequest(http.MethodPost, baseURL+"/projects/"+p.Id+"/stop", nil)
	if err != nil {
		return err
	}

	// Add access token as Bearer token in headers
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response is an error
	err = utils.Error(resp)
	if err != nil {
		return err
	}

	return nil
}
