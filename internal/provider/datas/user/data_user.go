package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

var dataSchema = map[string]*schema.Schema{
	"email": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
	"user_id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	},
}

func Data() *schema.Resource {

	return &schema.Resource{
		ReadContext: read,
		Schema:      dataSchema,
	}
}

func read(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	val, ok := data.GetOk("email")
	if ok {
		email := val.(string)
		user, err := client.GetUser(ctx, email)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Failed to fetch user",
				Detail:   err.Error(),
			})
			return diags
		}

		data.Set("email", user.Person.Email)
		data.Set("user_id", user.ID)
		data.Set("name", user.Name)
		data.SetId(user.ID)
	}

	return diags
}
