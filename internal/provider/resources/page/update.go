package page

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func update(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	if data.HasChange("title") {
		newTitle := data.Get("title").(string)
		err := client.UpdatePageTitle(ctx, data.Id(), newTitle)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldnt update title",
				Detail:   err.Error(),
			})
		}
	}

	if data.HasChange("title") {
		newTitle := data.Get("title").(string)
		err := client.UpdatePageTitle(ctx, data.Id(), newTitle)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldn't update title",
				Detail:   err.Error(),
			})
		}
	}

	if data.HasChange("parent_page_id") {

		parent := model.NewParentFromData(data)

		err := client.UpdatePageParent(ctx, data.Id(), parent)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldn't update page parent",
				Detail:   err.Error(),
			})
		}
	}

	if data.HasChange("database") {

		parent := model.NewParentFromData(data)

		err := client.UpdatePageParent(ctx, data.Id(), parent)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Couldn't update page parent",
				Detail:   err.Error(),
			})
		}
	}

	return diags
}
