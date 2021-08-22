package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			ResourcesMap:   map[string]*schema.Resource{},
			DataSourcesMap: map[string]*schema.Resource{},

			ConfigureContextFunc: providerConfigure,
		}
		return provider
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics

	return nil, diags
}
