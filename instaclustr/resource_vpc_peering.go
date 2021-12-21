package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVpcPeering() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpcPeeringCreate,
		Read:   resourceVpcPeeringRead,
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

			"peer_vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_account_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"peer_subnets"},
			},

			"peer_subnets": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:      true,
				ConflictsWith: []string{"peer_subnet"},
			},

			"peer_region": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"aws_vpc_connection_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVpcPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	cdcID, err := VpcPeeringCreate(d, meta)
	if err != nil {
		return fmt.Errorf("[Error] Error creating VPC peering request object: %s", err)
	}
	var createData CreateVPCPeeringRequest
	createData, err = createVpcPeeringRequest(d)
	if err != nil {
		return fmt.Errorf("[Error] Error creating VPC peering request: %s", err)
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating VPC peering request: %s", err)
	}
	var id string
	id, err = client.CreateVpcPeering(cdcID, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error creating cluster: %s", err)
	}
	d.SetId(id)
	d.Set("vpc_peering_id", id)
	d.Set("cdc_id", cdcID)
	var vpcPeering *VPCPeering
	vpcPeering, err = client.ReadVpcPeering(cdcID, id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}

	d.Set("aws_vpc_connection_id", vpcPeering.AWSVpcConnectionID)

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
	d.Set("aws_vpc_connection_id", vpcPeering.AWSVpcConnectionID)
	if len(vpcPeering.PeerSubnet) == 0 {
		d.Set("peer_subnets", vpcPeering.PeerSubnets)
	} else if len(vpcPeering.PeerSubnets) == 0 {
		d.Set("peer_subnet", vpcPeering.PeerSubnet)
	}
	d.Set("peer_region", vpcPeering.PeerRegion)

	log.Printf("[INFO] Fetched VPC peering %s info from the remote server.", vpcPeering.ID)
	return nil
}

func resourceVpcPeeringUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("[Error] The VPC peering connection doesn't support update")
}

func createVpcPeeringRequest(d *schema.ResourceData) (CreateVPCPeeringRequest, error) {
	result := CreateVPCPeeringRequest{
		PeerVpcID:     d.Get("peer_vpc_id").(string),
		PeerAccountID: d.Get("peer_account_id").(string),
		PeerRegion:    d.Get("peer_region").(string),
	}
	if _, isSet := d.GetOk("peer_subnet"); isSet {
		result.PeerSubnet = d.Get("peer_subnet").(string)
	} else if _, isSet := d.GetOk("peer_subnets"); isSet {
		result.PeerSubnets = d.Get("peer_subnets").(*schema.Set).List()
	} else {
		return result, fmt.Errorf("[Error] Error creating peering request - at least one subnet must be specified")
	}
	return result, nil
}
