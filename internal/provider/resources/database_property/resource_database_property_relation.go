package database_property

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func RelationResource() *schema.Resource {

	_schema := map[string]*schema.Schema{
		"database": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "ID of database this property belongs to",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "title of this property",
		},
		"related_database": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "ID of database this property relates to",
		},
	}

	createOrUpdate := createOrUpdate(func(data *schema.ResourceData) model.DatabaseProperty {
		name := data.Get("name").(string)
		relatedDatabaseId := data.Get("related_database").(string)
		return model.NewRelationDatabaseProperty(&name, relatedDatabaseId)
	})

	return &schema.Resource{
		ReadContext:   read,
		CreateContext: createOrUpdate,
		UpdateContext: createOrUpdate,
		DeleteContext: delete,
		Schema:        _schema,
	}
}
