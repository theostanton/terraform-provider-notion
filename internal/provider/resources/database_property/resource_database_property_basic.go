package database_property

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func BasicResource(propertyType string) *schema.Resource {

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
	}

	createOrUpdate := createOrUpdate(func(data *schema.ResourceData) model.DatabaseProperty {
		name := data.Get("name").(string)
		return model.NewBasicDatabaseProperty(&name, propertyType)
	})

	return &schema.Resource{
		ReadContext:   read,
		CreateContext: createOrUpdate,
		UpdateContext: createOrUpdate,
		DeleteContext: delete,
		Schema:        _schema,
	}
}
