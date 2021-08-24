package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceSchema = map[string]*schema.Schema{
	"title": {
		Type:     schema.TypeString,
		Required: true,
	},
	"parent": {
		Type:     schema.TypeString,
		Required: true,
	},
	"title_column_title": {
		Type:     schema.TypeString,
		Required: true,
	},
}

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: delete,
		Schema:        resourceSchema,
	}
}
