package database

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

func update(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	if data.HasChange("title") {
		newTitle := data.Get("title").(string)
		err := client.UpdateDatabaseTitle(ctx, data.Id(), newTitle)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt update title",
				Detail:   err.Error(),
			})
		}
	}

	if data.HasChange("title") {
		newTitle := data.Get("title").(string)
		err := client.UpdateDatabaseTitle(ctx, data.Id(), newTitle)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt update title",
				Detail:   err.Error(),
			})
		}
	}

	if data.HasChange("title_column_title") {
		newTitleColumnTitle := data.Get("title_column_title").(string)
		err := client.UpdateDatabaseTitleColumnTitle(ctx, data.Id(), newTitleColumnTitle)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt update title column title",
				Detail:   err.Error(),
			})
		}
	}

	return diags
}
