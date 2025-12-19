package mobula

import (
	"encoding/json"
	"fmt"
)

// APIError represents an error returned by the Mobula API
type APIError struct {
	StatusCode int
	Message    string
	RawBody    string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("mobula API error (status %d): %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("mobula API error (status %d): %s", e.StatusCode, e.RawBody)
}

// parseError attempts to parse an error response
func parseError(statusCode int, body []byte) error {
	apiErr := &APIError{
		StatusCode: statusCode,
		RawBody:    string(body),
	}

	// Try to parse error message from response
	var errResp struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(body, &errResp); err == nil {
		if errResp.Error != "" {
			apiErr.Message = errResp.Error
		} else if errResp.Message != "" {
			apiErr.Message = errResp.Message
		}
	}

	return apiErr
}
