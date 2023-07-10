package project

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/paastech-cloud/cli/pkg/utils"
)

type projectCreationRequest struct {
	Name string `json:"name"`
}

type projectCreationResponse struct {
	Content Project `json:"content"`
}

type Project struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Send a project creation request to the PaaSTech API and returns a Project object if successful
func CreateProject(baseURL string, accessToken string, name string) (Project, error) {
	// Create JSON request body
	request, err := json.Marshal(projectCreationRequest{
		Name: name,
	})
	if err != nil {
		return Project{}, err
	}

	// Create POST request to API
	req, err := http.NewRequest(http.MethodPost, baseURL+"/projects", bytes.NewReader(request))
	if err != nil {
		return Project{}, err
	}

	// Add access token as Bearer token in headers
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Project{}, err
	}
	defer resp.Body.Close()

	// Check if the response is an error
	err = utils.Error(resp)
	if err != nil {
		return Project{}, err
	}

	// Parse JSON body
	var res projectCreationResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return Project{}, err
	}

	return res.Content, nil
}

func (p *Project) Delete(baseURL string, accessToken string) error {
	// Create DELETE request to API
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/projects/"+p.Id, nil)
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

	// Parse JSON body
	var project Project
	err = json.NewDecoder(resp.Body).Decode(&project)
	if err != nil {
		return err
	}

	return nil
}
