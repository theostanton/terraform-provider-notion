package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	data_database "github.com/theostanton/terraform-provider-notion/internal/provider/datas/database"
	"github.com/theostanton/terraform-provider-notion/internal/provider/datas/database_entries"
	data_page "github.com/theostanton/terraform-provider-notion/internal/provider/datas/page"
	"github.com/theostanton/terraform-provider-notion/internal/provider/datas/user"
	"github.com/theostanton/terraform-provider-notion/internal/provider/resources/database"
	"github.com/theostanton/terraform-provider-notion/internal/provider/resources/database_property"
	"github.com/theostanton/terraform-provider-notion/internal/provider/resources/page"
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
				"notion_page":           page.PageResource(),
				"notion_database_entry": page.DatabasePageResource(),
				// todo API doesn't seem to handle workspace pages like it says
				//"notion_workspace_page":                     page.WorkspacePageResource(),
				"notion_database":                           database.Resource(),
				"notion_database_property_select":           database_property.SelectResource(),
				"notion_database_property_multi_select":     database_property.MultiSelectResource(),
				"notion_database_property_number":           database_property.NumberResource(),
				"notion_database_property_relation":         database_property.RelationResource(),
				"notion_database_property_rollup":           database_property.RollupResource(),
				"notion_database_property_rich_text":        database_property.BasicResource("rich_text"),
				"notion_database_property_date":             database_property.BasicResource("date"),
				"notion_database_property_people":           database_property.BasicResource("people"),
				"notion_database_property_checkbox":         database_property.BasicResource("checkbox"),
				"notion_database_property_url":              database_property.BasicResource("url"),
				"notion_database_property_email":            database_property.BasicResource("email"),
				"notion_database_property_created_time":     database_property.BasicResource("created_time"),
				"notion_database_property_created_by":       database_property.BasicResource("created_by"),
				"notion_database_property_last_edited_time": database_property.BasicResource("last_edited_time"),
				"notion_database_property_last_edited_by":   database_property.BasicResource("last_edited_by"),
			},

			DataSourcesMap: map[string]*schema.Resource{
				"notion_database":         data_database.Data(),
				"notion_user":             user.Data(),
				"notion_page":             data_page.Data(),
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

	_, err = client.GetCurrentUser(ctx)

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
