package database_property

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

func delete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	databaseId := data.Get("database").(string)

	databasePropertyId, err := client.DeleteDatabaseProperty(ctx, databaseId, data.Id())

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Delete Database Property API call failed",
			Detail:   err.Error(),
		})
		return diags
	}
	if databasePropertyId != "" {
		data.SetId("")
	}

	return diags
}
