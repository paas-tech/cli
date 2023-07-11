package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Structure of an error response from the API
type errorResponse struct {
	Message    interface{} `json:"message"`
	Error      string      `json:"error"`
	StatusCode int         `json:"statusCode"`
}

// Returns an error following the HTTP response code:
//
//	4xx: Gives the error messages from the response body
//	5xx: Gives a generic error message
//
// If the error code is not one of these, no error is returned
func Error(resp *http.Response) error {
	switch {
	// handle 4xx errors
	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		var respErr errorResponse
		err := json.NewDecoder(resp.Body).Decode(&respErr)
		if err != nil {
			return err
		}

		// switch case to fix messages being either an array or a string
		switch msg := respErr.Message.(type) {
		case string:
			return errors.New(msg)
		case []string:
			return errors.New(strings.Join(msg, ", "))
		}

		return errors.New("Unknown error")
	// handle 4xx errors
	case resp.StatusCode >= 500 && resp.StatusCode < 600:
		// handles 5xx errors
		return errors.New("Server error")

	// handle other status codes
	default:
		return nil
	}
}
