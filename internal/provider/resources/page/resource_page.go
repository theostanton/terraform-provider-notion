package page

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var pageResourceSchema = map[string]*schema.Schema{
	"title": {
		Type:     schema.TypeString,
		Required: true,
	},
	"parent_page_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"url": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func PageResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: deleteResource,
		Schema:        pageResourceSchema,
	}
}
