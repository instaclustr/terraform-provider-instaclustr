package instaclustr

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceKafkaUserList() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceKafkaUserListRead,

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

func dataSourceKafkaUserListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	usernameList, err := client.ReadKafkaUserList(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error fetching the kafka user list: %s", err)
	}

	d.SetId(fmt.Sprintf("%s-user-list", d.Get("cluster_id").(string)))
	d.Set("username_list", usernameList)

	log.Printf("[INFO] Fetched Kafka user list in %s.", d.Get("cluster_id").(string))
	return nil
}
