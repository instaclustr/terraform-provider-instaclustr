package instaclustr

import (
"fmt"
"log"

"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceKafkaSchemaRegistryUserList() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceKafkaSchemaRegistryUserListRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"username_list":  &schema.Schema {
				Type:     schema.TypeList,
				Elem:     &schema.Schema {
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceKafkaSchemaRegistryUserListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	usernameList, err := client.ReadKafkaSchemaRegistryUserList(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error fetching the kafka schema registry user list: %s", err)
	}

	d.SetId(fmt.Sprintf("%s-user-list", d.Get("cluster_id").(string)))
	d.Set("username_list", usernameList)

	log.Printf("[INFO] Fetched Kafka Schema Registry user list in %s.", d.Get("cluster_id").(string))
	return nil
}

