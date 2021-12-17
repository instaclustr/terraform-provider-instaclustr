package instaclustr

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVpcPeeringDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
	}
	cdcID := cluster.DataCentres[0].ID
	vpcPeeringID := d.Get("vpc_peering_id").(string)
	log.Printf("[INFO] Deleting VPC peering connection %s.", vpcPeeringID)
	err = client.DeleteVpcPeering(cdcID, vpcPeeringID)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting VPC peering connection: %s", err)
	}

	d.SetId("")
	d.Set("vpc_peering_id", "")
	d.Set("cdc_id", "")
	log.Printf("[INFO] VPC peering connection %s has been marked for deletion.", vpcPeeringID)
	return nil
}

func resourceVpcPeeringStateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "&")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("Unexpected format of ID (%q), expected <CLUSTER-ID>&<VPC-PEERING-ID>", d.Id())
	}

	d.Set("cluster_id", idParts[0])
	d.Set("vpc_peering_id", idParts[1])
	return []*schema.ResourceData{d}, nil
}
