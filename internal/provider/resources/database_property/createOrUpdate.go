package database_property

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func createOrUpdate(invoke func(data *schema.ResourceData) model.DatabaseProperty) func(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {

	return func(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {

		client := m.(*api.Client)
		var diags diag.Diagnostics

		var databaseProperty model.DatabaseProperty

		name := data.Get("name").(string)
		databaseId := data.Get("database").(string)

		var databasePropertyId string
		if data.Id() == "" {
			databasePropertyId = name
		} else {
			databasePropertyId = data.Id()
		}

		databaseProperty = invoke(data)

		databasePropertyId, err := client.SetDatabaseProperty(ctx, databaseId, databasePropertyId, databaseProperty)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Set Database Property API call failed",
				Detail:   err.Error(),
			})
			return diags
		}
		if databasePropertyId != "" {
			data.SetId(databasePropertyId)
		}

		return diags
	}

}