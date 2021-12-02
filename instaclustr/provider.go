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
			"api_hostname": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  DefaultApiHostname,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"instaclustr_cluster":        resourceCluster(),
			"instaclustr_encryption_key": resourceEncryptionKey(),
			"instaclustr_firewall_rule":  resourceFirewallRule(),
			"instaclustr_vpc_peering":    resourceVpcPeering(),
			"instaclustr_kafka_user":     resourceKafkaUser(),
			"instaclustr_kafka_topic":    resourceKafkaTopic(),
			"instaclustr_kafka_acl":      resourceKafkaAcl(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"instaclustr_cluster_credentials": dataSourceClusterCredentials(),
			"instaclustr_kafka_user_list":     dataSourceKafkaUserList(),
			"instaclustr_kafka_topic_list":    dataSourceKafkaTopicList(),
			"instaclustr_kafka_acl_list":      dataSourceKafkaAclList(),
		},
	}
	provider.ConfigureFunc = providerConfigure

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:          d.Get("username").(string),
		ApiKey:            d.Get("api_key").(string),
		apiServerHostname: d.Get("api_hostname").(string),
	}

	config.Init()

	return &config, nil
}
