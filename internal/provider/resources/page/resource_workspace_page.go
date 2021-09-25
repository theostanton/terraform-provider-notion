package page

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var workspacePageResourceSchema = map[string]*schema.Schema{
	"title": {
		Type:     schema.TypeString,
		Required: true,
	},
	"url": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func WorkspacePageResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: deleteResource,
		Schema:        workspacePageResourceSchema,
	}
}
