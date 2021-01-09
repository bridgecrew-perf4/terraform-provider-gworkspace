package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/micahkemp/terraform-provider-gworkspace/gworkspace"
)

func main() {
	opts := &plugin.ServeOpts{
		ProviderFunc: gworkspace.Provider,
	}

	plugin.Serve(opts)
}
