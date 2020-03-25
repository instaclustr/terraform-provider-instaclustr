package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceEncryptionKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceEncryptionKeyCreate,
		Read:   resourceEncryptionKeyRead,
		Update: resourceEncryptionKeyUpdate,
		Delete: resourceEncryptionKeyDelete,

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alias": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEncryptionKeyCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Adding encryption key.")
	client := meta.(*Config).Client

	createData := EncryptionKey{
		Alias: d.Get("alias").(string),
		ARN:   d.Get("arn").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error adding encryption key: %s", err)
	}

	id, err := client.CreateEncryptionKey(jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error adding encryption key: %s", err)
	}
	d.SetId(id)
	d.Set("key_id", id)
	log.Printf("[INFO] Encyption key %s has been added.", id)
	return nil
}

func resourceEncryptionKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("key_id").(string)
	log.Printf("[INFO] Reading encryption key %s.", id)
	keyResource, err := client.ReadEncryptionKey(id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading encryption key: %s", err)
	}

	d.SetId(keyResource.ID)
	d.Set("key_id", keyResource.ID)
	d.Set("alias", keyResource.Alias)
	d.Set("arn", keyResource.ARN)
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
	err := client.DeleteEncryptionKey(id)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting encryption key: %s", err)
	}

	d.SetId("")
	d.Set("key_id", "")
	log.Printf("[INFO] Encryption key %s has been deleted.", id)
	return nil
}
