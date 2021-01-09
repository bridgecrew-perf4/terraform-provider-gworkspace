package gworkspace

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credential_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path to JSON-formatted Service Account credentials file",
			},
			"impersonate_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address of the account to impersonate",
			},
		},
	}

	return provider
}
