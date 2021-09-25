package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"net/http"
)

type searchFilter struct {
	Value    string `json:"value"`
	Property string `json:"property"`
}
type searchRequest struct {
	Query  string       `json:"query"`
	Filter searchFilter `json:"filter"`
}

type searchPagesResult struct {
	Pages []model.Page `json:"results"`
}

type searchDatabasesResult struct {
	Databases []model.Database `json:"results"`
}

func (client *Client) SearchPages(ctx context.Context, query string) ([]model.Page, error) {

	body := searchRequest{
		Query: query,
		Filter: searchFilter{
			Value:    "page",
			Property: "object",
		},
	}

	req, err := client.generatePost(ctx, "search", body)
	if err != nil {
		return nil, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to search pages: %w", parseErrorResponse(res))
	}

	var response *searchPagesResult
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return response.Pages, nil
}

func (client *Client) SearchDatabases(ctx context.Context, query string) ([]model.Database, error) {

	body := searchRequest{
		Query: query,
		Filter: searchFilter{
			Value:    "database",
			Property: "object",
		},
	}

	req, err := client.generatePost(ctx, "search", body)
	if err != nil {
		return nil, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to search databases: %w", parseErrorResponse(res))
	}

	var response *searchDatabasesResult
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return response.Databases, nil
}
