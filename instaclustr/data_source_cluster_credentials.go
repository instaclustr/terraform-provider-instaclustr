package instaclustr

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceClusterCredentials() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceClusterCredentialsRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     	schema.TypeString,
				Required: 	true,
			},
			"cluster_password": {
				Type:		schema.TypeString,
				Computed:	true,
				Sensitive:	true,
			},
			"certificate_download": {
				Type:		schema.TypeString,
				Computed:	true,
			},
		},
	}
}

func dataSourceClusterCredentialsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)

	log.Printf("[INFO] Reading credentials of cluster %s.", id)
	cluster, err := client.ReadCluster(id)

	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster credentials: #{err}")
	}

	d.SetId(fmt.Sprintf("%s-credentials", id))
	d.Set("cluster_id", id)
	d.Set("cluster_password", cluster.InstaclustrUserPassword)
	d.Set("certificate_download", cluster.ClusterCertificateDownload)

	return nil
}
