package auth_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paastech-cloud/cli/pkg/auth"
)

func TestLogin(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method and path
		if r.Method != http.MethodPost || r.URL.Path != "/auth/login" {
			t.Errorf("Expected POST request to /auth/login, got %s request to %s", r.Method, r.URL.Path)
		}

		// Check the request body
		expectedBody := `{"email":"test@example.com","password":"password123"}`
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error reading request body: %s", err)
		}
		if string(body) != expectedBody {
			t.Errorf("Expected request body %s, got %s", expectedBody, string(body))
		}

		// Send a mock response
		response := `{"status":"OK","content":{"accessToken":"mockAccessToken"}}`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Set the base URL to the mock server URL
	baseURL := server.URL

	// Call the Login function
	accessToken, err := auth.Login(baseURL, "test@example.com", "password123")
	if err != nil {
		t.Errorf("Error calling Login: %s", err)
	}

	// Check the returned access token
	expectedAccessToken := "mockAccessToken"
	if accessToken != expectedAccessToken {
		t.Errorf("Expected access token %s, got %s", expectedAccessToken, accessToken)
	}
}
