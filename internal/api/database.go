package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"net/http"
)

type GetDatabaseResponse struct {
	Results []model.Database `json:"results"`
	HasMore bool             `json:"has_more"`
}

func (client *Client) GetDatabase(ctx context.Context, databaseId string) (database model.Database, err error) {
	path := fmt.Sprintf("databases/%s", databaseId)
	req, err := client.generateGet(ctx, path)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return model.Database{}, fmt.Errorf("failed to find database: %w", parseErrorResponse(res))
	}

	var response *GetDatabaseResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return model.Database{}, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return response.Results[0], nil
}

type QueryDatabaseResponse struct {
	Results []model.Page `json:"results"`
	HasMore bool         `json:"has_more"`
}

func (client *Client) QueryDatabase(ctx context.Context, databaseId string) (pages []model.Page, err error) {
	path := fmt.Sprintf("databases/%s/query", databaseId)
	type QueryDatabaseBody struct {
	}
	req, err := client.generatePost(ctx, path, QueryDatabaseBody{})
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to find database: %w", parseErrorResponse(res))
	}

	var response *QueryDatabaseResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return response.Results, nil
}

func (client *Client) CreateDatabase(ctx context.Context, database model.Database) (id string, err error) {

	req, err := client.generatePost(ctx, "databases", database)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to create database: %w", parseErrorResponse(res))
	}

	var response *model.Database

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return *response.Id, nil
}

func (client *Client) SetDatabaseProperty(ctx context.Context, databaseId string, databasePropertyId string, databaseProperty model.DatabaseProperty) (id string, err error) {

	type SetDatabasePropertyBody struct {
		Properties map[string]model.DatabaseProperty `json:"properties"`
	}

	body := SetDatabasePropertyBody{
		Properties: map[string]model.DatabaseProperty{
			databasePropertyId: databaseProperty,
		},
	}

	path := fmt.Sprintf("databases/%s", databaseId)
	req, err := client.generatePatch(ctx, path, body)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to edit database property: %w", parseErrorResponse(res))
	}

	var response *model.Database

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	if databaseProperty.Name != nil {
		return *databaseProperty.Name, nil
	} else {
		return databasePropertyId, nil
	}
}

func (client *Client) DeleteDatabaseProperty(ctx context.Context, databaseId string, databasePropertyId string) (id string, err error) {

	type DeleteDatabasePropertyBody struct {
		Properties map[string]*model.DatabaseProperty `json:"properties"`
	}

	body := DeleteDatabasePropertyBody{
		Properties: map[string]*model.DatabaseProperty{
			databasePropertyId: nil,
		},
	}

	path := fmt.Sprintf("databases/%s", databaseId)
	req, err := client.generatePatch(ctx, path, body)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to delete database property: %w", parseErrorResponse(res))
	}

	var response *model.Database

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return databasePropertyId, nil
}

func (client *Client) UpdateDatabaseTitle(ctx context.Context, databaseId string, title string) (err error) {

	path := fmt.Sprintf("databases/%s", databaseId)

	body := model.Database{
		Title: []model.RichText{
			model.NewRichText(title),
		},
	}

	req, err := client.generatePatch(ctx, path, body)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update database title: %w", parseErrorResponse(res))
	}

	return nil
}

func (client *Client) UpdateDatabaseTitleColumnTitle(ctx context.Context, databaseId string, titleColumnTitle string) (err error) {

	path := fmt.Sprintf("databases/%s", databaseId)

	body := model.Database{
		Properties: map[string]model.DatabaseProperty{
			"title": model.NewTitleDatabaseProperty(titleColumnTitle),
		},
	}

	req, err := client.generatePatch(ctx, path, body)
	if err != nil {
		return
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update database title: %w", parseErrorResponse(res))
	}

	return nil
}
