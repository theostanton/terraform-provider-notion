package database_property

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"github.com/theostanton/terraform-provider-notion/internal/utils/logger"
	"strings"
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
			logger.Info("databasePropertyId=%s from name", name)
			databasePropertyId = name
		} else {
			logger.Info("databasePropertyId=%s from id", data.Id())
			databasePropertyId = data.Id()
		}

		databaseProperty = invoke(data)

		_databasePropertyId := ""
		attemptsLeft := 3
		for _databasePropertyId == "" {
			var err error
			_databasePropertyId, err = client.SetDatabaseProperty(ctx, databaseId, databasePropertyId, databaseProperty)

			switch {
			case err == nil:
				databasePropertyId = _databasePropertyId
			case strings.Contains(err.Error(), "Conflict occurred while saving") && attemptsLeft > 0:
				attemptsLeft = attemptsLeft - 1
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  fmt.Sprintf("Conflict error on creating database, retrying %d more times", attemptsLeft),
					Detail:   err.Error(),
				})
			default:
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Set Database Property API call failed",
					Detail:   err.Error(),
				})
				return diags
			}

		}

		if databasePropertyId != "" {
			data.SetId(databasePropertyId)
		}

		return diags
	}

}
