package data_database

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/utils/logger"
)

var dataSchema = map[string]*schema.Schema{
	"query": {
		Type:     schema.TypeString,
		Required: true,
	},
	"title": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"url": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func Data() *schema.Resource {

	return &schema.Resource{
		ReadContext: read,
		Schema:      dataSchema,
	}
}

func read(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	client := m.(*api.Client)

	query := data.Get("query").(string)
	databases, err := client.SearchDatabases(ctx, query)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed tosearch pages from API",
			Detail:   err.Error(),
		})
		return
	}

	if len(databases) == 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to find databases from API for query",
		})
		return
	}

	if len(databases) > 1 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Found multiple databases from API for query",
		})
		return
	}

	database := databases[0]

	logger.InfoObject("database", database)

	title := database.ExtractTitle()
	err = data.Set("title", title)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to set title from API response",
			Detail:   err.Error(),
		})
	}

	err = data.Set("url", database.Url)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to set URL from API response",
			Detail:   err.Error(),
		})
	}

	data.SetId(*database.Id)

	return
}
