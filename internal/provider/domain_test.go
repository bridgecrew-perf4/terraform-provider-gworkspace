package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceDomain(t *testing.T) {
	//t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDomain,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"gworkspace_domain.foo", "name", regexp.MustCompile("^example.com$")),
				),
			},
		},
	})
}

const testAccResourceDomain = `
resource "gworkspace_domain" "foo" {
  name = "example.com"
}
`
