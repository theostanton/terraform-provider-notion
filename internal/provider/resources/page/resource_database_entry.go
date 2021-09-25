package page

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var databasePageResourceSchema = map[string]*schema.Schema{
	"title": {
		Type:     schema.TypeString,
		Required: true,
	},
	"database": {
		Type:     schema.TypeString,
		Required: true,
	},
	"url": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func DatabasePageResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: archive,
		Schema:        databasePageResourceSchema,
	}
}
