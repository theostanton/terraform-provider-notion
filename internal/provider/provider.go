package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/provider/datas/database_entries"
	"github.com/theostanton/terraform-provider-notion/internal/provider/datas/user"
	"github.com/theostanton/terraform-provider-notion/internal/provider/resources/database"
	"github.com/theostanton/terraform-provider-notion/internal/provider/resources/database_property"
)

func New() func() *schema.Provider {
	return func() *schema.Provider {
		provider := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"token": {
					Type:      schema.TypeString,
					Required:  true,
					Sensitive: true,
								},
			},
			ResourcesMap: map[string]*schema.Resource{
				"notion_database":                    database.Resource(),
				"notion_database_property_select":    database_property.SelectResource(),
				"notion_database_property_number":    database_property.NumberResource(),
				"notion_database_property_rich_text": database_property.RichTextResource(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"notion_user":             user.Data(),
				"notion_database_entries": database_entries.Data(),
			},

			ConfigureContextFunc: providerConfigure,
		}
		return provider
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	tokenVal, ok := d.GetOk("token")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No token",
			Detail:   "Notion token is required",
		})

		return nil, diags

	}

	token := tokenVal.(string)

	if "" == token {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid token",
			Detail:   "Empty string is not a valid token",
		})
		return nil, diags
	}

	client, err := api.NewClient(token)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to create Notion client",
			Detail:   err.Error(),
		})
		return nil, diags
	}

	_, err = client.ListUsers(ctx)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to test client",
			Detail:   err.Error(),
		})
		return nil, diags

	}

	return client, diags
}
