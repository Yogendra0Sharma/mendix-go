package mendix

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const url = ""

func newTransport() *http.Client {
	return &http.Client{}
}

// Client is Mendix GPT-3 API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	username   string
	apiKey     string
}

// NewClient creates new Mendix API client with PATS.
func NewClientWithPATKey(apiKey string) *Client {
	return &Client{
		BaseURL:    url,
		HTTPClient: newTransport(),
		apiKey:     apiKey,
	}
}

// NewClient creates new Mendix API client with PATS.
func NewClientWithUserNameAndAPIKey(username string, apiKey string) *Client {
	return &Client{
		BaseURL:    url,
		HTTPClient: newTransport(),
		username:   username,
		apiKey:     apiKey,
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	if c.username != "" {
		req.Header.Set("Mendix-Username", c.username)
		req.Header.Set("Mendix-ApiKey", c.apiKey)
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("MxToken %s", c.apiKey))
	}

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil || errRes.Error == nil {
			return fmt.Errorf("error, status code: %d", res.StatusCode)
		}
		return fmt.Errorf("error, status code: %d, message: %s", res.StatusCode, errRes.Error.Message)
	}

	if v != nil {
		if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, suffix)
}
