package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSiteVerificationToken(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSiteVerificationToken,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						// this is a pretty inflexible test, and should be made more robust
						"data.gworkspace_siteverification_token.foo", "token", regexp.MustCompile("^google-site-verification="),
					),
				),
			},
		},
	})
}

const testAccResourceSiteVerificationToken = `
resource "gworkspace_domain" "foo" {
  name = "example.com"
}

data "gworkspace_siteverification_token" "foo" {
  site_id = "example.com"
  type = "INET_DOMAIN"
  method = "DNS"
}
`
