package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

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
				Optional: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Optional: true,
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
	if err != nil {
		return fmt.Errorf("[Error] Error adding encryption key: %s", err)
	}

	id, err := client.EncryptionKeyAdd(jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error adding encryption key: %s", err)
	}
	d.SetId(after(id, "%v"))
	d.Set("key_id", after(id, "%v"))
	log.Printf("[INFO] Encyption key %s has been added.", after(id, "%v"))
	return nil
}

func after(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func resourceEncryptionKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("key_id").(string)
	log.Printf("[INFO] Reading encryption key %s.", id)
	keyResource, err := client.EncryptionKeyRead(id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}

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
	err := client.EncryptionKeyDelete(id)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting encryption key: %s", err)
	}

	d.SetId("")
	d.Set("key_id", "")
	log.Printf("[INFO] Encryption key %s has been deleted.", id)
	return nil
}
