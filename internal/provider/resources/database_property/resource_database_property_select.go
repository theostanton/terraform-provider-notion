package database_property

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func SelectResource() *schema.Resource {

	_schema := map[string]*schema.Schema{
		"database": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "ID of database this property belongs to",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "Name",
			Description: "title of this property",
		},
		"options": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "Sorted list of options available for this property.",
			ValidateDiagFunc: func(value interface{}, path cty.Path) (diags diag.Diagnostics) {

				itemsRaw := value.(map[string]interface{})

				validColours := []string{"default", "gray", "brown", "orange", "yellow", "green", "blue", "purple", "pink", "red"}

				for title, colour := range itemsRaw {
					valid := false
					for _, validColour := range validColours {
						if colour.(string) == validColour {
							valid = true
						}
					}
					if !valid {
						diags = append(diags, diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Invalid colour property",
							Detail:   fmt.Sprintf("'%s' is not a valid colour for option %s", colour, title),
						})
					}
				}

				return
			},
		},
	}

	createOrUpdate := createOrUpdate(func(data *schema.ResourceData) model.DatabaseProperty {
		var name string
		if data.Id() == "" {
			name = data.Get("name").(string)
		}
		rawSelectOptions := data.Get("options").(map[string]interface{})
		if rawSelectOptions != nil {
			options := make([]model.SelectOption, len(rawSelectOptions))
			for optionName, rawSelectOption := range rawSelectOptions {
				options = append(options, model.SelectOption{
					Name:  optionName,
					Color: rawSelectOption.(string),
				})
			}
			return model.NewSelectDatabaseProperty(&name, &options)
		} else {
			return model.NewSelectDatabaseProperty(&name, nil)
		}
	})
	return &schema.Resource{
		ReadContext:   read,
		CreateContext: createOrUpdate,
		UpdateContext: createOrUpdate,
		DeleteContext: delete,
		Schema:        _schema,
	}
}
