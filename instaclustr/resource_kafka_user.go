package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKafkaUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaUserCreate,
		Read:   resourceKafkaUserRead,
		Update: resourceKafkaUserUpdate,
		Delete: resourceKafkaUserDelete,

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

func resourceKafkaUserCreate(d *schema.ResourceData, meta interface{}) error {
	cluster_id := d.Get("cluster_id").(string)
	username := d.Get("username").(string)

	log.Printf("[INFO] Creating Kafka user in %s.", cluster_id)
	client := meta.(*Config).Client

	// Cluster has to reach running state first
	cluster, err := client.ReadCluster(cluster_id)
	if err != nil {
		return fmt.Errorf("[Error] Error in getting the status of the cluster: %s", err)
	}
	if cluster.ClusterStatus != "RUNNING" {
		return fmt.Errorf("[Error] Cluster %s is not RUNNING.", cluster_id)
	}
		
	// just use linear search for now to check if the user going to be created is already in the system
	usernameList, err := client.ReadKafkaUserList(cluster_id)
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving kafka user list: %s", err)
	}
	for _, str := range usernameList {
		if str == username {
			// user is already set, so we don't change anything
			d.SetId(fmt.Sprintf("%s-%s", cluster_id, username))
			log.Printf("[INFO] Kafka user %d already exists in %s.", username, cluster_id)
			return nil
		}
	}

	createData := CreateKafkaUserRequest{
		Username:              username,
		Password:              d.Get("password").(string),
		InitialPermissions:    d.Get("initial_permissions").(string),
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka user creation request: %s", err)
	}

	err = client.CreateKafkaUser(cluster_id, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka user: %s", err)
	}
	
	d.SetId(fmt.Sprintf("%s-%s", cluster_id, username))

	log.Printf("[INFO] Kafka user %s has been created.", username)
	return nil
}

func resourceKafkaUserRead(d *schema.ResourceData, meta interface{}) error {
	// there is almost no point in reading, because the API only returns the username
	return nil
}

func resourceKafkaUserUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Changing Kafka user password in %s.", d.Get("cluster_id").(string))
	client := meta.(*Config).Client

	createData := UpdateKafkaUserRequest{
		Username:              d.Get("username").(string),
		Password:              d.Get("password").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka user update request: %s", err)
	}

	err = client.UpdateKafkaUser(d.Get("cluster_id").(string), jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error updating the password for kafka user: %s", err)
	}

	log.Printf("[INFO] The password for Kafka user %s has been updated.", d.Get("username").(string))
	return nil
}

func removeKafkaUserResource(d *schema.ResourceData) {
	d.SetId("")
	d.Set("cluster_id", "")
	d.Set("username", "")
	d.Set("password", "")
	d.Set("initial_permissions", "")
}

func resourceKafkaUserDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting Kafka user %s in %s.", d.Get("username").(string), d.Get("cluster_id"))
	client := meta.(*Config).Client

	createData := DeleteKafkaUserRequest{
		Username:              d.Get("username").(string),
	}
	
	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka user update request: %s", err)
	}
	
	err = client.DeleteKafkaUser(d.Get("cluster_id").(string), jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting Kafka user: %s", err)
	}
	
	removeKafkaUserResource(d)

	log.Printf("[INFO] Kafka user %s has been deleted.", d.Get("username").(string))
	return nil
}
