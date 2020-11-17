package instaclustr

import (
"encoding/json"
"fmt"
"log"

"github.com/hashicorp/terraform/helper/schema"
)

func resourceKafkaRestProxyUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceKafkaRestProxyUserRead,
		Update: resourceKafkaRestProxyUserUpdate,
		Delete: resourceKafkaRestProxyUserDelete,

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

func resourceKafkaRestProxyUserRead(d *schema.ResourceData, meta interface{}) error {
	// there is almost no point in reading, because the API only returns the username
	return nil
}

func resourceKafkaRestProxyUserUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Changing Kafka Rest Proxy user password in %s.", d.Get("cluster_id").(string))
	client := meta.(*Config).Client

	createData := UpdateKafkaRestProxyUserRequest{
		Username:              d.Get("username").(string),
		Password:              d.Get("password").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka rest proxy user update request: %s", err)
	}

	err = client.UpdateKafkaRestProxyUser(d.Get("cluster_id").(string), jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error updating the password for kafka rest proxy user: %s", err)
	}

	log.Printf("[INFO] The password for Kafka Rest Proxy user %s has been updated.", d.Get("username").(string))
	return nil
}

func resourceKafkaRestProxyUserDelete(d *schema.ResourceData, meta interface{}) error {
	// there is almost no point in reading, because the API only returns the username
	return nil
}