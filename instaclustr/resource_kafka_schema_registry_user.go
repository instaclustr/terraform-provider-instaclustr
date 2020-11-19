package instaclustr

import (
"encoding/json"
"fmt"
"log"

"github.com/hashicorp/terraform/helper/schema"
)

func resourceKafkaSchemaRegistry() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaSchemaRegistryUserCreate,
		Read:   resourceKafkaSchemaRegistryUserRead,
		Update: resourceKafkaSchemaRegistryUserUpdate,
		Delete: resourceKafkaSchemaRegistryUserDelete,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"password": {
				Type:     schema.TypeString,
				Required: true,
			},

			"initial_permissions": {
				Type:     schema.TypeString,
				Optional: true,
				Default: "none",
				ForceNew: true,
			},
		},
	}
}

func resourceKafkaSchemaRegistryUserCreate(d *schema.ResourceData, meta interface{}) error {
	// we do not support create user feature for kafka schema registry at the moment
	return nil
}

func resourceKafkaSchemaRegistryUserRead(d *schema.ResourceData, meta interface{}) error {
	// there is almost no point in reading, because the API only returns the username
	return nil
}

func resourceKafkaSchemaRegistryUserUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Changing Kafka Schema Registry user password in %s.", d.Get("cluster_id").(string))
	client := meta.(*Config).Client

	createData := UpdateKafkaSchemaRegistryUserRequest{
		Username:              d.Get("username").(string),
		Password:              d.Get("password").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka schema registry user update request: %s", err)
	}

	err = client.UpdateKafkaSchemaRegistryUser(d.Get("cluster_id").(string), jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error updating the password for kafka schema registry user: %s", err)
	}

	log.Printf("[INFO] The password for Kafka Schema Registry user %s has been updated.", d.Get("username").(string))
	return nil
}

func resourceKafkaSchemaRegistryUserDelete(d *schema.ResourceData, meta interface{}) error {
	// there is almost no point in reading, because the API only returns the username
	return nil
}