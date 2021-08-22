package api

import (
	"context"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"net/http"
	"os"
	"testing"
)

func TestClient_CreateDatabase(t *testing.T) {

	parentPageId := "0f69b08d4b3c4974bf2c2a021280cf97"
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
					"Name":      model.NewTitleDatabaseProperty(nil),
					"Some text": model.NewRichTextDatabaseProperty(),
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
