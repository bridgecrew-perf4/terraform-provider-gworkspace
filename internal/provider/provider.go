package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	jsonCredsFileKey    = "json_credentials_file"
	jsonCredsFileEnv	= "GWORKSPACE_JSON_CREDS"
	impersonateEmailKey = "impersonate_email"
	impersonateEmailEnv = "GWORKSPACE_IMPERSONATE"
	customerKey         = "customer"
	customerEnv         = "GWORKSPACE_CUSTOMER"

)

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent

		client := client{
			jsonCredentialsFile: d.Get(jsonCredsFileKey).(string),
			impersonateEmail: d.Get(impersonateEmailKey).(string),
			customerNumber: d.Get(customerKey).(string),
		}
		return &client, nil
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
				"gworkspace_domain": resourceDomain(),
			},
			Schema: map[string]*schema.Schema{
				jsonCredsFileKey: &schema.Schema{
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(jsonCredsFileEnv, nil),
					Description: "JSON file of Google Workspace credentials",
				},
				impersonateEmailKey: &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(impersonateEmailEnv, nil),
					Description: "Email address to impersonate",
				},
				customerKey: {
					Type: 		 schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc(customerEnv, nil),
					Description: "Customer number",
				},
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}
