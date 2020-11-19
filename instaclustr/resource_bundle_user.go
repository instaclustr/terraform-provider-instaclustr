package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBundleUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceBundleUserCreate,
		Read:   resourceBundleUserRead,
		Update: resourceBundleUserUpdate,
		Delete: resourceBundleUserDelete,

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

			"bundle_name": {
				Type: 	  schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBundleUserCreate(d *schema.ResourceData, meta interface{}) error {
	cluster_id := d.Get("cluster_id").(string)
	username := d.Get("username").(string)
	bundle_name := d.Get("bundle_name").(string)

	log.Printf("[INFO] Creating %s user in %s.", bundle_name,cluster_id)
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
	usernameList, err := client.ReadBundleUserList(cluster_id, bundle_name)
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving kafka user list: %s", err)
	}
	for _, str := range usernameList {
		if str == username {
			// user is already set, so we don't change anything
			d.SetId(fmt.Sprintf("%s-%s", cluster_id, username))
			log.Printf("[INFO] %s user %d already exists in %s.", bundle_name, username, cluster_id)
			return nil
		}
	}

	createData := CreateBundleUserRequest{
		Username:              username,
		Password:              d.Get("password").(string),
		InitialPermissions:    d.Get("initial_permissions").(string),
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating %s user creation request: %s", bundle_name, err)
	}

	err = client.CreateKafkaUser(cluster_id, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error creating %s user: %s", bundle_name, err)
	}

	d.SetId(fmt.Sprintf("%s-%s", cluster_id, username))

	log.Printf("[INFO] %s user %s has been created.", bundle_name, username)
	return nil
}

func resourceBundleUserRead(d *schema.ResourceData, meta interface{}) error {
	// there is almost no point in reading, because the API only returns the username
	return nil
}

func resourceBundleUserUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Changing %s user password in %s.", d.Get("bundle_name").(string), d.Get("cluster_id").(string))
	client := meta.(*Config).Client

	createData := UpdateBundleUserRequest{
		Username:              d.Get("username").(string),
		Password:              d.Get("password").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating %s user update request: %s", d.Get("bundle_name").(string), err)
	}

	err = client.UpdateBundleUser(d.Get("cluster_id").(string), d.Get("bundle_name").(string), jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error updating the password for %s user: %s", d.Get("bundle_name").(string), err)
	}

	log.Printf("[INFO] The password for %s user %s has been updated.", d.Get("bundle_name"), d.Get("username").(string))
	return nil
}

func removeBundleUserResource(d *schema.ResourceData) {
	d.SetId("")
	d.Set("cluster_id", "")
	d.Set("username", "")
	d.Set("password", "")
	d.Set("initial_permissions", "")
}

func resourceBundleUserDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting %s user %s in %s.", d.Get("bundle_name"), d.Get("username").(string), d.Get("cluster_id"))
	client := meta.(*Config).Client

	createData := DeleteBundleUserRequest{
		Username:              d.Get("username").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating %s user update request: %s", d.Get("bundle_name"), err)
	}

	err = client.DeleteBundleUser(d.Get("cluster_id").(string), d.Get("bundle_name"), jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting %s user: %s", d.Get("bundle_name"), err)
	}

	removeBundleUserResource(d)

	log.Printf("[INFO] %s user %s has been deleted.", d.Get("bundle_name"), d.Get("username").(string))
	return nil
}