package main

import (
	"github.com/dikhan/terraform-provider-openapi/v2/openapi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
	"os"
)

// Generate docs for website
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var ProviderHost = "https://api.instaclustr.com"

const ProviderName = "instaclustr"

func main() {

	swaggerHostOverride := os.Getenv("IC_TF_PROVIDER_SWAGGER_HOST_OVERRIDE") //needed to target locally run API to generate docs for testing/preview

	finalProviderHost := ProviderHost
	if len(swaggerHostOverride) > 0 {
		finalProviderHost = swaggerHostOverride
	}

	providerOpenAPIURL := finalProviderHost + "/provisioning/v2/swagger-for-terraform.yaml"

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
