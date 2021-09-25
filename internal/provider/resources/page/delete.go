package page

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

func deleteResource(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)

	pageId := data.Id()

	err := client.ArchivePage(ctx, pageId)

	var diags diag.Diagnostics

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Couldn't delete page",
			Detail:   err.Error(),
		})
	}

	return diags
}
