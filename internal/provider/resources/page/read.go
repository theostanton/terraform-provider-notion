package page

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

func read(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	client := m.(*api.Client)

	page, err := client.GetPage(ctx, data.Id())

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to Get Page from API",
			Detail:   err.Error(),
		})
		return
	}

	title, err := page.ExtractPageTitle()
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Couldn't extract title from page",
			Detail:   err.Error(),
		})
	} else {
		err = data.Set("title", title)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Failed to set title from API response",
				Detail:   err.Error(),
			})
		}
	}

	if page.Parent.PageId != "" {
		err = data.Set("parent_page_id", page.Parent.PageId)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Failed to set parent page from API response",
				Detail:   err.Error(),
			})
		}
	}

	if page.Parent.DatabaseId != "" {
		err = data.Set("database", page.Parent.DatabaseId)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Failed to set database parent from API response",
				Detail:   err.Error(),
			})
		}
	}

	err = data.Set("url", page.Url)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to set url from API response",
			Detail:   err.Error(),
		})
	}

	return diags
}
