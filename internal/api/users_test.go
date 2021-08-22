package api

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestClient_ListUsers(t *testing.T) {

	validToken := os.Getenv("NOTION_TOKEN")
	ctx := context.Background()

	if validToken == "" {
		t.Error("Need validToken value")
		return
	}

	type fields struct {
		apiToken string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantUsersLength int
		wantErr         bool
	}{
		{
			name:            "Valid token",
			fields:          fields{apiToken: validToken},
			args:            args{ctx: ctx},
			wantUsersLength: 3,
			wantErr:         false,
		},
		{
			name:    "Invalid token",
			fields:  fields{apiToken: ""},
			args:    args{ctx: ctx},
			wantErr: true,
		},
		{
			name:    "Expired token",
			fields:  fields{apiToken: ""},
			args:    args{ctx: ctx},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				apiToken:   tt.fields.apiToken,
				httpClient: http.DefaultClient,
			}
			gotUsers, err := client.ListUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantUsersLength != len(gotUsers) {
				t.Errorf("ListUsers() len(gotUsers) = %v, wantUsersLength %v", gotUsers, tt.wantUsersLength)
			}
		})
	}
}
