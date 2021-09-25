package api

import (
	"context"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"net/http"
	"os"
	"testing"
)

func TestClient_CreatePage(t *testing.T) {

	apiToken, ok := os.LookupEnv("NOTION_TOKEN")
	if ok == false {
		t.Error("NOTION_TOKEN is required")
		return
	}
	testParentPageId,ok := os.LookupEnv("NOTION_TEST_PARENT_PAGE_ID")
	if ok == false {
		t.Error("NOTION_TEST_PARENT_PAGE_ID is required")
		return
	}
	testDatabaseId,ok := os.LookupEnv("NOTION_TEST_DATABASE_ID")
	if ok == false {
		t.Error("NOTION_TEST_DATABASE_ID is required")
		return
	}
	ctx := context.Background()

	tests := []struct {
		name    string
		page    model.PagePatch
		wantErr bool
	}{
		{
			name:    "Basic Page",
			page:    model.NewPage("Some Basic Page", testParentPageId),
			wantErr: false,
		},
		{
			name:    "Basic Workspace Page",
			page:    model.NewWorkspacePage("Some Basic Workspace Page"),
			wantErr: true, // Despite docs, doesn't look like API can handle this
		},
		{
			name:    "Basic Database Entry Page",
			page:    model.NewDatabaseEntryPage("Some Basic Database Entry Page", testDatabaseId),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				apiToken:   apiToken,
				httpClient: http.DefaultClient,
			}
			gotPage, err := client.CreatePage(ctx, tt.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if  tt.wantErr==false  && gotPage.Id == nil {
				t.Errorf("CreatePage() error gotPage.Id == nil")
				return
			}

		})
	}
}
