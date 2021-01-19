package provider

import (
	"context"
	admin "google.golang.org/api/admin/directory/v1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	domainNameKey = "name"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Description: "Domain associated with Google Workspace",

		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		DeleteContext: resourceDomainDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			domainNameKey: {
				Description: "Domain name",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	service, err := c.newAdminServiceWithScopes(ctx, admin.AdminDirectoryDomainScope)
	if err != nil {
		return diag.Errorf("unable to get newAdminServiceWithScopes: %s", err)
	}

	insertDomains := &admin.Domains{
		DomainName: d.Get(domainNameKey).(string),
	}

	insertedDomains, err := service.Domains.Insert(c.customerNumber, insertDomains).Do()
	if err != nil {
		return diag.Errorf("failure inserting domains: %s", err)
	}

	d.SetId(insertedDomains.DomainName)

	return diag.Diagnostics{}
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	service, err := c.newAdminServiceWithScopes(ctx, admin.AdminDirectoryDomainScope)
	if err != nil {
		return diag.Errorf("unable to get newAdminServiceWithScopes: %s", err)
	}

	_, err = service.Domains.Get(c.customerNumber, d.Id()).Do()
	if err != nil {
		return diag.Errorf("unable to get domain %s: %s", d.Id(), err)
	}

	return diag.Diagnostics{}
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c:= meta.(*client)

	service, err := c.newAdminServiceWithScopes(ctx, admin.AdminDirectoryDomainScope)
	if err != nil {
		return diag.Errorf("unable to get newAdminServiceWithScopes: %s", err)
	}

	err = service.Domains.Delete(c.customerNumber, d.Id()).Do()
	if err != nil {
		return diag.Errorf("unable to delete domain %s: %s", d.Id(), err)
	}

	return diag.Diagnostics{}
}
