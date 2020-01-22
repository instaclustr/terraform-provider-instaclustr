package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVpcPeering() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpcPeeringCreate,
		Read:   resourceVpcPeeringRead,
		Update: resourceVpcPeeringUpdate,
		Delete: resourceVpcPeeringDelete,

		Schema: map[string]*schema.Schema{
			"vpc_peering_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"cdc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"peer_vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVpcPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating VPC peering request.")
	client := meta.(*Config).Client

	var cdcID string
	for {
		cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
		if err != nil {
			return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
		}
		if cluster.ClusterStatus == "PROVISIONED" {
			cdcID = cluster.DataCentres[0].ID
			break
		}
		time.Sleep(5 * time.Second)
	}

	createData := CreateVPCPeeringRequest{
		PeerVpcID:     d.Get("peer_vpc_id").(string),
		PeerAccountID: d.Get("peer_account_id").(string),
		PeerSubnet:    d.Get("peer_subnet").(string),
		PeerRegion:    d.Get("peer_region").(string),
	}

	var jsonStr []byte
	jsonStr, err := json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating VPC peering request: %s", err)
	}

	id, err := client.CreateVpcPeering(cdcID, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error creating cluster: %s", err)
	}
	d.SetId(id)
	d.Set("vpc_peering_id", id)
	d.Set("cdc_id", cdcID)
	log.Printf("[INFO] VPC peering request %s has been created.", id)
	return nil
}

func resourceVpcPeeringRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
	}
	cdcID := cluster.DataCentres[0].ID
	vpcPeeringID := d.Get("vpc_peering_id").(string)

	log.Printf("[INFO] Reading the status of VPC peering connection %s.", vpcPeeringID)
	vpcPeering, err := client.ReadVpcPeering(cdcID, vpcPeeringID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}

	d.SetId(vpcPeering.ID)
	d.Set("vpc_peering_id", vpcPeering.ID)
	d.Set("cdc_id", vpcPeering.ClusterDataCentre)
	d.Set("peer_vpc_id", vpcPeering.PeerVpcID)
	d.Set("peer_account_id", vpcPeering.PeerAccountID)

	before, _ := d.GetChange("peer_subnet")
	d.Set("peer_subnet", before)
	d.Set("peer_region", vpcPeering.PeerRegion)

	log.Printf("[INFO] Fetched VPC peering %s info from the remote server.", vpcPeering.ID)
	return nil
}

func resourceVpcPeeringUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("[Error] The VPC peering connection doesn't support update")
}

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
