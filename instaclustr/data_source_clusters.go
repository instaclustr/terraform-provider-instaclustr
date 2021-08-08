package instaclustr

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceClustersList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClustersListRead,

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cluster_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cassandra_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"running_node_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"derived_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sla_tier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pci_compliance": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceClustersListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	d.SetId("instaclustr-clusters")

	log.Print("[INFO] Listing all active clusters")
	clusterList, err := client.ListClusters()
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster credentials: #{err}")
	}

	clusters := make([]map[string]interface{}, 0)

	for _, c := range *clusterList {
		cluster := make(map[string]interface{})
		cluster["cluster_id"] = c.ID
		cluster["cluster_name"] = c.Name
		cluster["cassandra_version"] = c.CassandraVersion
		cluster["node_count"] = c.NodeCount
		cluster["running_node_count"] = c.RunningNodeCount
		cluster["derived_status"] = c.DerivedStatus
		cluster["sla_tier"] = c.SlaTier
		cluster["pci_compliance"] = c.PciCompliance
		clusters = append(clusters, cluster)
	}
	d.Set("cluster", clusters)
	return nil
}
