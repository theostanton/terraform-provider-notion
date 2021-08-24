package database_property

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func read(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	client := m.(*api.Client)
	databaseId := data.Get("database").(string)

	database, err := client.GetDatabase(ctx, databaseId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Get Database API call failed",
			Detail:   err.Error(),
		})
		return diags
	}

	name := data.Get("name").(string)

	var property model.DatabaseProperty
	foundProperty := false
	for _, databaseProperty := range database.Properties {
		if *databaseProperty.Name == name {
			property = databaseProperty
			foundProperty = true
		}
	}

	if foundProperty == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Couldnt find property",
			Detail:   fmt.Sprintf("Couldnt find property for name=%s in Get Database response", name),
		})
		return diags
	}

	switch *property.Type {
	case "rich_text":
		return
	case "select":
		if property.Select == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt find select property option",
				Detail:   fmt.Sprintf("Couldnt find select info on property for name=%s in Get Database response", name),
			})
			return
		}
		options := map[string]string{}
		for _, option := range *property.Select.Options {
			options[option.Name] = option.Color
		}
		err = data.Set("options", options)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt set options value",
				Detail:   err.Error(),
			})
			return
		}

		return
	case "number":
		if property.Number == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt find number property option",
				Detail:   fmt.Sprintf("Couldnt find number info on property for name=%s in Get Database response", name),
			})
			return
		}

		err = data.Set("format", *property.Number.Format)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt set format value",
				Detail:   err.Error(),
			})
			return
		}
	default:
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Couldnt handle option type",
			Detail:   fmt.Sprintf("Unknown option type=%s", property.Type),
		})
		return

	}
	return
}
