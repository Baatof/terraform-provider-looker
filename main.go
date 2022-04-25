package main

import (
	"github.com/adeo/terraform-provider-looker/pkg/looker"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: looker.Provider,
	})
}
