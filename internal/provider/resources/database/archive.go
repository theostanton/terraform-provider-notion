package database

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

func archive(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if false {
		client := m.(*api.Client)

		databaseId := data.Id()

		err := client.ArchivePage(ctx, databaseId)

		var diags diag.Diagnostics

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Couldn't delete page",
				Detail:   err.Error(),
			})
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Couldn't delete database",
			Detail:   fmt.Sprintf("Doesnt look like Notion API can archive databases yet, go do it yourself at https://notion.so/%s", data.Id()),
		})
	}

	return diags
}
