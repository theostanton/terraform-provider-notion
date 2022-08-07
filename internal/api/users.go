package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/theostanton/terraform-provider-notion/internal/model"
)

type ListUsersResponse struct {
	Results []model.User `json:"results"`
	HasMore bool         `json:"has_more"`
}

func (client *Client) ListUsers(ctx context.Context) (users []model.User, err error) {
	req, err := client.generateGet(ctx, "users")
	if err != nil {
		return nil, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to find user: %w", parseErrorResponse(res))
	}

	var response *ListUsersResponse

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return response.Results, nil
}

func (client *Client) GetUser(ctx context.Context, email string) (user model.User, err error) {
	users, err := client.ListUsers(ctx)
	if err != nil {
		return model.User{}, err
	}

	for _, user := range users {
		if user.Person.Email == email {
			return user, nil
		}
	}

	return model.User{}, errors.New("Cannot find user")
}

func (client *Client) GetCurrentUser(ctx context.Context) (user model.User, err error) {
	none := model.User{}
	req, err := client.generateGet(ctx, "users/me")
	if err != nil {
		return none, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return none, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return none, fmt.Errorf("failed to find current user: %w", parseErrorResponse(res))
	}

	var response *model.User

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return none, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return *response, nil
}
