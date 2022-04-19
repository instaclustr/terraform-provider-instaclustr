package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGCPVpcPeering() *schema.Resource {
	return &schema.Resource{
		Create: GCPresourceVpcPeeringCreate,
		Read:   GCPresourceVpcPeeringRead,
		Update: resourceVpcPeeringUpdate,
		Delete: resourceVpcPeeringDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVpcPeeringStateImport,
		},

		Schema: map[string]*schema.Schema{

			"vpc_peering_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cdc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"peer_vpc_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_project_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_subnets": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func GCPresourceVpcPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	cdcID, err := VpcPeeringCreate(d, meta)
	if err != nil {
		return fmt.Errorf("[Error] Error creating VPC peering request object: %s", err)
	}
	var createData CreateGCPVPCPeeringRequest
	createData, err = GCPcreateVpcPeeringRequest(d)
	if err != nil {
		return fmt.Errorf("[Error] Error creating GCP VPC peering request: %s", err)
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating GCP VPC peering request: %s", err)
	}
	var id string
	id, err = client.CreateVpcPeering(cdcID, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error in creating the cluster: %s", err)
	}
	d.SetId(id)
	d.Set("vpc_peering_id", id)
	d.Set("cdc_id", cdcID)
	var vpcPeering *GCPVPCPeering
	vpcPeering, err = client.GCPReadVpcPeering(cdcID, id)
	if err != nil {
		return fmt.Errorf("[Error] Error in reading GCP VPC peering connection: %s", err)
	}

	d.Set("peer_project_id", vpcPeering.PeerProjectID)

	log.Printf("[INFO] VPC peering request %s has been created.", id)
	return nil
}

func GCPresourceVpcPeeringRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
	}
	cdcID := cluster.DataCentres[0].ID
	vpcPeeringID := d.Get("vpc_peering_id").(string)

	log.Printf("[INFO] Reading the status of VPC peering connection %s.", vpcPeeringID)
	vpcPeering, err := client.GCPReadVpcPeering(cdcID, vpcPeeringID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}
	err = MapGCPVPCPeeringToResource(d, vpcPeering)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}
	return nil

}

func MapGCPVPCPeeringToResource(d *schema.ResourceData, vpcPeering *GCPVPCPeering) error {
	if vpcPeering.ID == "" {
		return nil
	}
	d.SetId(vpcPeering.ID)
	d.Set("vpc_peering_id", vpcPeering.ID)
	d.Set("cdc_id", vpcPeering.ClusterDataCentre)
	d.Set("peer_vpc_network_name", vpcPeering.PeerVPCNetworkName)
	d.Set("peer_subnets", vpcPeering.PeerSubnets)
	d.Set("peer_project_id", vpcPeering.PeerProjectID)

	log.Printf("[INFO] Fetched VPC peering %s info from the remote server.", vpcPeering.ID)
	return nil
}

func resourceGCPVpcPeeringUpdate(d *schema.ResourceData) error {
	return fmt.Errorf("[Error] The VPC peering connection doesn't support update")

}

func GCPcreateVpcPeeringRequest(d *schema.ResourceData) (CreateGCPVPCPeeringRequest, error) {
	result := CreateGCPVPCPeeringRequest{
		PeerVPCNetworkName: d.Get("peer_vpc_network_name").(string),
		PeerProjectID:      d.Get("peer_project_id").(string),
	}
	if _, isSet := d.GetOk("peer_subnets"); isSet {
		result.PeerSubnets = d.Get("peer_subnets").(*schema.Set).List()
	} else {
		return result, fmt.Errorf("[Error] Error creating GCP VPC peering request - Please check the subnets atleast one subnet must be specified")
	}
	return result, nil
}
