package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	baseUrl    = "https://api.notion.com/v1"
	apiVersion = "2021-08-16"
)

// Client is used for HTTP requests to the Notion API.
type Client struct {
	apiToken   string
	httpClient *http.Client
}

func NewClient(apiToken string) (*Client, error) {

	if apiToken == "" {
		return nil, errors.New("Empty string is not a valid token")
	}

	client := &Client{
		apiToken:   apiToken,
		httpClient: http.DefaultClient,
	}

	return client, nil
}

func (client *Client) generateRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", baseUrl, path)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", client.apiToken))
	req.Header.Set("Notion-Version", apiVersion)
	req.Header.Set("User-Agent", "terraform-provider-notion")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (client *Client) generateGet(ctx context.Context, path string) (*http.Request, error) {
	return client.generateRequest(ctx, "GET", path, nil)
}

func (client *Client) generatePost(ctx context.Context, path string, body interface{}) (*http.Request, error) {
	bodyBytes, err := json.Marshal(body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(bodyBytes)
	return client.generateRequest(ctx, "POST", path, bodyReader)
}

func (client *Client) generatePatch(ctx context.Context, path string, body interface{}) (*http.Request, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(bodyBytes)
	return client.generateRequest(ctx, "PATCH", path, bodyReader)
}
