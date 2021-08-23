package api

import (
	"context"
	"fmt"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"net/http"
	"os"
	"testing"
)

const parentPageId = "0f69b08d4b3c4974bf2c2a021280cf97"

func TestClient_CreateDatabase(t *testing.T) {

	token := os.Getenv("NOTION_TOKEN")
	ctx := context.Background()

	if token == "" {
		t.Error("Need validToken value")
		return
	}

	tests := []struct {
		name     string
		database model.Database
		wantErr  bool
	}{
		{
			name: "Basic",
			database: model.NewDatabase(
				"Basic",
				model.NewParentFromPageId(parentPageId),
				map[string]model.DatabaseProperty{
					"title":     model.NewTitleDatabaseProperty("Name"),
					"Some text": model.NewRichTextDatabaseProperty(nil),
				},
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				apiToken:   token,
				httpClient: http.DefaultClient,
			}
			_, err := client.CreateDatabase(ctx, tt.database)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_SetDatabaseProperty(t *testing.T) {

	token := os.Getenv("NOTION_TOKEN")
	ctx := context.Background()

	if token == "" {
		t.Error("Need validToken value")
		return
	}

	client := &Client{
		apiToken:   token,
		httpClient: http.DefaultClient,
	}

	name := "Name"
	database := model.NewDatabase("Test", model.NewParentFromPageId(parentPageId), map[string]model.DatabaseProperty{
		"title": model.NewTitleDatabaseProperty(name),
	})
	databaseId, err := client.CreateDatabase(ctx, database)

	if err != nil {
		t.Error(fmt.Sprintf("failed to create test database err=%s", err.Error()))
		return
	}

	editedPropertyTitle := "New property edited"

	type args struct {
		databaseId         string
		databasePropertyId string
		databaseProperty   model.DatabaseProperty
	}
	tests := []struct {
		name    string
		args    args
		wantId  string
		wantErr bool
	}{
		{
			name: "New property",
			args: args{
				databaseId:         databaseId,
				databasePropertyId: "New property",
				databaseProperty:   model.NewRichTextDatabaseProperty(nil),
			},
			wantId:  "New property",
			wantErr: false,
		},
		{
			name: "Edit property",
			args: args{
				databaseId:         databaseId,
				databasePropertyId: "New property",
				databaseProperty:   model.NewRichTextDatabaseProperty(&editedPropertyTitle),
			},
			wantId:  editedPropertyTitle,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := client.SetDatabaseProperty(ctx, tt.args.databaseId, tt.args.databasePropertyId, tt.args.databaseProperty)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetDatabaseProperty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("SetDatabaseProperty() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
