package main

import (
	"github.com/dikhan/terraform-provider-openapi/v3/openapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
	"os"
)

// Generate docs for website
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

const ProviderName = "instaclustr"

func main() {

	providerHost := os.Getenv("IC_API_URL") //needed to target locally run API to generate docs for testing/preview

	if len(providerHost) == 0 {
		providerHost = "https://api.instaclustr.com"
	}

	providerOpenAPIURL := providerHost + "/cluster-management/v2/swagger-for-terraform.yaml"

	p := openapi.ProviderOpenAPI{ProviderName: ProviderName}
	serviceProviderConfig := &openapi.ServiceConfigV1{
		SwaggerURL: providerOpenAPIURL,
	}

	provider, err := p.CreateSchemaProviderFromServiceConfiguration(serviceProviderConfig)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize the terraform provider: %s", err)
	}

	plugin.Serve(
		&plugin.ServeOpts{
			ProviderFunc: func() *schema.Provider {
				return provider
			},
		},
	)
}
