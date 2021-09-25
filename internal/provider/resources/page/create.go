package page

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/model"
	"strings"
)

func create(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	title, ok := data.GetOk("title")
	if ok == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No title provided",
			Detail:   "title is required for creating a database",
		})
		return diags
	}

	parent := model.NewParentFromData(data)

	pagePatch := model.NewPage(title.(string), parent)

	var storedPage model.Page
	done := false
	attemptsLeft := 3
	for done == false {
		_storedPage, err := client.CreatePage(ctx, pagePatch)

		switch {
		case err == nil:
			done = true
			storedPage = _storedPage
		case strings.Contains(err.Error(), "Conflict occurred while saving.") && attemptsLeft > 0:
			attemptsLeft = attemptsLeft - 1
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("Conflict error on creating page, retrying %d more times", attemptsLeft),
				Detail:   err.Error(),
			})
		default:
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Create Page API call failed",
				Detail:   err.Error(),
			})
			return diags
		}

	}

	err := data.Set("url", storedPage.Url)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to set url",
			Detail:   err.Error(),
		})
	}

	data.SetId(*storedPage.Id)

	return diags
}
