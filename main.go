package main

import (
	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return instaclustr.Provider()
		},
	})
}
