package instaclustr

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceKafkaTopicList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKafkaTopicListRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"topics": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceKafkaTopicListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	topicList, err := client.ReadKafkaTopicList(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error fetching the kafka user list: %s", err)
	}

	d.SetId(fmt.Sprintf("%s-topic-list", d.Get("cluster_id").(string)))
	d.Set("topics", topicList.Topics)

	log.Printf("[INFO] Fetched Kafka topic list in %s.", d.Get("cluster_id").(string))
	return nil
}
