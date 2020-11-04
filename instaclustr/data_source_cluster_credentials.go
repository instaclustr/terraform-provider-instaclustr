package instaclustr

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func dataSourceClusterCredentials() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceClusterCredentialsRead,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
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
	clusterCredentials, err := client.ReadClusterCredentials(id)

	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster credentials: #{err}")
	}

	d.SetId(fmt.Sprintf("%s-credentials", id))
	d.Set("cluster_id", id)
	d.Set("cluster_password", clusterCredentials.ClusterPassword)
	d.Set("cluster_certificate_download", clusterCredentials.ClusterCertificateDownload)

	return nil
}
