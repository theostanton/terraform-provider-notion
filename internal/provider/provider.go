package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/provider/datas/user"
	"github.com/theostanton/terraform-provider-notion/internal/provider/resources/database"
)

func New() func() *schema.Provider {
	return func() *schema.Provider {
		provider := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"token": &schema.Schema{
					Type:      schema.TypeString,
					Required:  true,
					Sensitive: true,
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"notion_database": database.Resource(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"notion_user": user.Data(),
			},

			ConfigureContextFunc: providerConfigure,
		}
		return provider
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics

	return nil, diags
}
