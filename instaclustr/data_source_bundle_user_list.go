package instaclustr

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceBundleUserList() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceBundleUserListRead,

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

			"bundle_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceBundleUserListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	usernameList, err := client.ReadBundleUserList(d.Get("cluster_id").(string), d.Get("bundle_name").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error fetching the %s user list: %s", d.Get("bundle_name"), err)
	}

	d.SetId(fmt.Sprintf("%s-user-list", d.Get("cluster_id").(string)))
	d.Set("username_list", usernameList)

	log.Printf("[INFO] Fetched %s user list in %s.", d.Get("bundle_name"), d.Get("cluster_id").(string))
	return nil
}