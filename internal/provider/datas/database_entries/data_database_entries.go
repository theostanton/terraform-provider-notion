package database_entries

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dataSchema = map[string]*schema.Schema{
	"database": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
	"entries": &schema.Schema{
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeMap,
		},
		Optional: true,
	},
}

func Data() *schema.Resource {

	return &schema.Resource{
		ReadContext: read,
		Schema:      dataSchema,
	}
}
