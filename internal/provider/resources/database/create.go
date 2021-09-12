package database

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"strings"
)

func create(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	abort := false
	database := model.Database{}

	title, ok := data.GetOk("title")
	if ok {
		database.Title = []model.RichText{
			model.NewRichText(title.(string)),
		}
	} else {
		abort = true
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No title provided",
			Detail:   "title is required for creating a database",
		})
	}

	titleColumnTitle := data.Get("title_column_title").(string)
	if ok {
		database.Properties = map[string]model.DatabaseProperty{
			titleColumnTitle: model.NewTitleDatabaseProperty(titleColumnTitle),
		}
	} else {
		abort = true
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No title column title provided",
			Detail:   "title column title is required for creating a database",
		})
	}

	parent, ok := data.GetOk("parent")
	if ok {
		parent := model.NewParentFromPageId(parent.(string))
		database.Parent = &parent
	} else {
		abort = true
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No parent provided",
			Detail:   "parent is required for creating a database",
		})
	}

	if abort {
		return diags
	}

	databaseId := ""
	attemptsLeft := 3
	for databaseId == "" {
		_databaseId, err := client.CreateDatabase(ctx, database)

		switch {
		case err == nil:
			databaseId = _databaseId
		case strings.Contains(err.Error(), "Conflict occurred while saving.") && attemptsLeft > 0:
			attemptsLeft = attemptsLeft - 1
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("Conflict error on creating database, retrying %d more times", attemptsLeft),
				Detail:   err.Error(),
			})
		default:
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Create Database API call failed",
				Detail:   err.Error(),
			})
			return diags
		}

	}

	url := fmt.Sprintf("notion.so/%s", databaseId)
	err := data.Set("url", url)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to set url",
			Detail:   err.Error(),
		})
	}

	data.SetId(databaseId)

	return diags
}
