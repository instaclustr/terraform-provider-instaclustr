package instaclustr

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceKafkaAclList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKafkaAclListRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"acls": &schema.Schema {
				Type: schema.TypeList,
				Elem: &schema.Schema {
					Type:     schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func doDataSourceKafkaAclListRead(d KafkaAclResourceDataInterface, client KafkaAclAPIClientInterface) error {
	cluster_id := d.Get("cluster_id").(string)

	data := KafkaAcl {
		ResourceType:	"ANY",
		Operation: 	"ANY",
		PermissionType:	"ANY",
		PatternType: 	"ANY",
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("[Error] Error creating kafka ACL read request: %w", err)
	}
	acls, err := client.ReadKafkaAcls(cluster_id, jsonStr)

	var aclsString []string
	for _, acl := range acls {
		aclsString = append(aclsString, fmt.Sprintf("(principal=%s, host=%s, resourceType=%s, resourceName=%s, operation=%s, permissionType=%s, patternType=%s)", 
			acl.Principal, acl.Host, acl.ResourceType, acl.ResourceName, acl.Operation, acl.PermissionType, acl.PatternType))
	}

	d.SetId(fmt.Sprintf("%s-acl-list", cluster_id))
	d.Set("acls", aclsString)

	log.Printf("[INFO] Fetched Kafka acl list in %s.", cluster_id)
	return nil
}

// it's a bit ugly of an ugly hack to enable unit testing
func dataSourceKafkaAclListRead(d *schema.ResourceData, meta interface{}) error {
	return doDataSourceKafkaAclListRead(d, meta.(*Config).Client)
}

