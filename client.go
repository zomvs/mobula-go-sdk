package mobula

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	DefaultBaseURL = "https://api.mobula.io"
	DemoBaseURL    = "https://demo-api.mobula.io"
	DefaultTimeout = 15 * time.Second
)

// Client is the main Mobula API client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// Config holds the configuration for the Mobula client
type Config struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
	Timeout    time.Duration
}

// NewClient creates a new Mobula API client
func NewClient(config *Config) *Client {
	if config == nil {
		config = &Config{}
	}

	baseURL := config.BaseURL
	if baseURL == "" {
		if config.APIKey == "" {
			baseURL = DemoBaseURL
		} else {
			baseURL = DefaultBaseURL
		}
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		timeout := config.Timeout
		if timeout == 0 {
			timeout = DefaultTimeout
		}
		httpClient = &http.Client{
			Timeout: timeout,
		}
	}

	client := &Client{
		baseURL:    baseURL,
		apiKey:     config.APIKey,
		httpClient: httpClient,
	}

	return client
}

// doRequest performs an HTTP request
func (c *Client) doRequest(ctx context.Context, method, path string, queryParams url.Values, body interface{}) ([]byte, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}
	u.Path = path
	if queryParams != nil {
		u.RawQuery = queryParams.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.apiKey != "" {
		req.Header.Set("Authorization", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, parseError(resp.StatusCode, respBody)
	}

	return respBody, nil
}

// get performs a GET request
func (c *Client) get(ctx context.Context, path string, queryParams url.Values, result interface{}) error {
	respBody, err := c.doRequest(ctx, http.MethodGet, path, queryParams, nil)
	if err != nil {
		return err
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// post performs a POST request
func (c *Client) post(ctx context.Context, path string, body interface{}, result interface{}) error {
	respBody, err := c.doRequest(ctx, http.MethodPost, path, nil, body)
	if err != nil {
		return err
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
