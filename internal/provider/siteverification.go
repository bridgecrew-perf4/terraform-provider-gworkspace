package provider

import (
	"context"
	"fmt"
	"google.golang.org/api/siteverification/v1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	siteIDKey = "site_id"
	typeKey   = "type"
	methodKey = "method"
	tokenKey  = "token"
)

func resourceSiteVerificationToken() *schema.Resource {
	return &schema.Resource{
		Description: "Get Site Verification info",

		ReadContext: resourceSiteVerificationTokenRead,

		Schema: map[string]*schema.Schema{
			siteIDKey: {
				Description: "ID of the site being validated",
				Type:        schema.TypeString,
				Required:    true,
			},
			typeKey: {
				Description: "Type of site being validated",
				Type:        schema.TypeString,
				Required:    true,
			},
			methodKey: {
				Description: "The method of validation to use",
				Type:        schema.TypeString,
				Required:    true,
			},
			tokenKey: {
				Description: "The token that can be used for validation",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceSiteVerificationTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	service, err := c.newSiteVerificationServiceWithScopes(ctx, siteverification.SiteverificationScope)
	if err != nil {
		return diag.Errorf("unable to get newSiteVerificationServiceWithScopes: %s", err)
	}

	tokenRequestSite := siteverification.SiteVerificationWebResourceGettokenRequestSite{
		Identifier: d.Get(siteIDKey).(string),
		Type:       d.Get(typeKey).(string),
	}

	tokenRequest := siteverification.SiteVerificationWebResourceGettokenRequest{
		Site:               &tokenRequestSite,
		VerificationMethod: d.Get(methodKey).(string),
	}

	tokenResponse, err := service.WebResource.GetToken(&tokenRequest).Do()
	if err != nil {
		return diag.Errorf("unable to get Site Verification Token: %s", err)
	}

	d.SetId(idForSiteVerificationToken(d))
	d.Set(tokenKey, tokenResponse.Token)

	return diag.Diagnostics{}
}

func idForSiteVerificationToken(d *schema.ResourceData) string {
	return fmt.Sprintf("%s-%s-%s", d.Get(siteIDKey).(string), d.Get(typeKey).(string), d.Get(methodKey).(string))
}
