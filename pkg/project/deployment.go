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

type logsResponse struct {
	Content struct {
		Logs string `json:"logs"`
	} `json:"content"`
}

type statusRequest struct {
	ContainerNames []string `json:"container_names"`
}

type containerStatus struct {
	Name   string `json:"container_name"`
	Status string `json:"container_status"`
}
type statusResponse struct {
	Content struct {
		ContainerStatuses []containerStatus `json:"container_statuses"`
	} `json:"content"`
}

var status = map[string]string{
	"STATUS_RUNNING":  "🟢 Running",
	"STATUS_STOPPED":  "🛑 Stopped",
	"STATUS_STOPPING": "🛑 Stopping",
	"STATUS_STARTING": "🏁 Starting",
	"STATUS_UNKNOWN":  "👽 Unknown",
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

// Get project logs from PaaSTech
func (p *Project) Logs(baseURL string, accessToken string) (string, error) {
	// Create GET request to API
	req, err := http.NewRequest(http.MethodGet, baseURL+"/projects/"+p.Id+"/logs", nil)
	if err != nil {
		return "", err
	}

	// Add access token as Bearer token in headers
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Make request
	resp, err := http.DefaultClient.Do(req)
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
	var res logsResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}

	return res.Content.Logs, nil
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

// Get project status from PaaSTech
func (p *Project) Status(baseURL string, accessToken string) (string, error) {
	// Create JSON request body
	request, err := json.Marshal(statusRequest{
		ContainerNames: []string{p.Id},
	})
	if err != nil {
		return "", err
	}

	// Create POST request to API
	req, err := http.NewRequest(http.MethodPost, baseURL+"/projects/status", bytes.NewReader(request))
	if err != nil {
		return "", err
	}

	// Add access token as Bearer token in headers
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Make request
	resp, err := http.DefaultClient.Do(req)
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
	var res statusResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}

	return status[res.Content.ContainerStatuses[0].Status], nil
}
