package database_entries

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

func read(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics
	var err error

	databaseId := data.Get("database").(string)

	pages, err := client.QueryDatabase(ctx, databaseId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Couldnt fetch database",
			Detail:   fmt.Sprintf("Query Database API call failed err=%s", err.Error()),
		})
		return diags
	}

	var entries []map[string]interface{}
	for _, page := range pages {

		if page.Properties != nil {
			for _, property := range *page.Properties {
				if property.Type == "title" {
					titleItems := *property.Title
					var title string
					if len(titleItems) > 0 {
						title = titleItems[0].Text.Content
					} else {
						title = "NA"
					}
					entries = append(entries, map[string]interface{}{
						"Title": title,
					})
				}
			}
		}
	}

	err = data.Set("entries", entries)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Couldnt set entries",
			Detail:   err.Error(),
		})
		return diags
	}

	data.SetId(fmt.Sprintf("%s-entries", databaseId))
	return diags
}
