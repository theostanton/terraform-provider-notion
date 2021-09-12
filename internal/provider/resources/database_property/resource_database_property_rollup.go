package database_property

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func RollupResource() *schema.Resource {

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
		"relation_property": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the relation property this property is responsible for rolling up. This relation is in the same database where the new rollup property is being created.",
		},
		"rollup_property": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the property in the related database that is used as an input to function.",
		},
		"function": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The function that is evaluated for every page in the relation of the rollup.\nPossible values include: count_all, count_values, count_unique_values, count_empty, count_not_empty, percent_empty, percent_not_empty, sum, average, median, min, max, range, show_original",
		},
	}

	createOrUpdate := createOrUpdate(func(data *schema.ResourceData) model.DatabaseProperty {
		name := data.Get("name").(string)
		relationProperty := data.Get("relation_property").(string)
		rollupProperty := data.Get("rollup_property").(string)
		function := model.RollupFunction(data.Get("function").(string))
		return model.NewRollupDatabaseProperty(&name, rollupProperty, relationProperty, function)
	})

	return &schema.Resource{
		ReadContext:   read,
		CreateContext: createOrUpdate,
		UpdateContext: createOrUpdate,
		DeleteContext: delete,
		Schema:        _schema,
	}
}
