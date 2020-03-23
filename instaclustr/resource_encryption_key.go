package instaclustr

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceEncryptionKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceEncryptionKeyAdd,
		Read:   resourceEncryptionKeyRead,
		Update: resourceEncryptionKeyUpdate,
		Delete: resourceEncryptionKeyDelete,

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEncryptionKeyAdd(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Adding encryption key.")
	client := meta.(*Config).Client

	createData := EncryptionKey{
		Alias: d.Get("alias").(string),
		ARN:   d.Get("arn").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	log.Printf("[EBS TEST]: %s", jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error adding encryption key: %s", err)
	}

	id, err := client.EncryptionKeyAdd(jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error adding encryption key: %s", err)
	}
	d.SetId(id)
	d.Set("key_id", id)
	log.Printf("[INFO] Encyption key %s has been added.", id)
	return nil
}

func getResourceByID(resources *[]EncryptionKey, id string) (interface{}, error) {
	for _, resource := range *resources {
		if resource.ID == id {
			return resource, nil
		}
	}
	return nil, errors.New(id)
}

func resourceEncryptionKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("key_id").(string)
	log.Printf("[INFO] Reading encryption key %s.", id)
	keys, err := client.EncryptionKeyRead()
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}

	keyResource, err := getResourceByID(keys, id)
	if err != nil {
		return fmt.Errorf("[Error] Error encryption key %s does not exist", id)
	}

	d.Set("key_id", keyResource.(EncryptionKey).ID)
	d.Set("alias", keyResource.(EncryptionKey).Alias)
	d.Set("arn", keyResource.(EncryptionKey).ARN)
	log.Printf("[INFO] Read encyption key %s.", id)
	return nil
}

func resourceEncryptionKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("[Error] The encryption keys don't support update")
}

func resourceEncryptionKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("key_id").(string)
	log.Printf("[INFO] Deleting encryption key %s.", id)
	err := client.EncryptionKeyDelete(id)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting encryption key: %s", err)
	}

	d.SetId("")
	d.Set("key_id", "")
	log.Printf("[INFO] Encryption key %s has been deleted.", id)
	return nil
}
