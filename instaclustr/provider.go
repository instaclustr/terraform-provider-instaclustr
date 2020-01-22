package instaclustr

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_key": {
				Type:     schema.TypeString,
				Required: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"instaclustr_cluster":       resourceCluster(),
			"instaclustr_firewall_rule": resourceFirewallRule(),
			"instaclustr_vpc_peering":   resourceVpcPeering(),
		},
	}
	provider.ConfigureFunc = providerConfigure

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:     d.Get("username").(string),
		ApiKey:       d.Get("api_key").(string),
	}

	config.Init()

	return &config, nil
}
