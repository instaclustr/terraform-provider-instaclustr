package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAzureVpcPeering() *schema.Resource {
	return &schema.Resource{
		Create: AzureresourceVpcPeeringCreate,
		Read:   AzureresourceVpcPeeringRead,
		Update: resourceVpcPeeringUpdate,
		Delete: resourceVpcPeeringDelete,

		Importer: &schema.ResourceImporter{
			State: resourceVpcPeeringStateImport,
		},

		Schema: map[string]*schema.Schema{

			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cdc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"peer_resource_group": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_suscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"peer_vpc_net": {
				Type:     schema.TypeString,
				Required: true,
			},

			"resource_group": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"suscription_id": {
				Type:     schema.TypeString,
				Computed: true,
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

			"v_net": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func AzureresourceVpcPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	cdcID, err := VpcPeeringCreate(d, meta)
	if err != nil {
		return fmt.Errorf("[Error] Error creating VPC peering request object: %s", err)
	}
	var createData CreateAzureVPCPeeringRequest
	createData, err = AzurecreateVpcPeeringRequest(d)
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
	d.Set("peer_suscription_id", id)
	d.Set("cdc_id", cdcID)
	var vpcPeering *AzureVPCPeering
	vpcPeering, err = client.AzureReadVpcPeering(cdcID, id)
	if err != nil {
		return fmt.Errorf("[Error] Error in reading GCP VPC peering connection: %s", err)
	}

	d.Set("peer_suscription_id", vpcPeering.PeerSubscriptionId)

	log.Printf("[INFO] VPC peering request %s has been created.", id)
	return nil
}

func AzureresourceVpcPeeringRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
	}
	cdcID := cluster.DataCentres[0].ID
	vpcPeeringID := d.Get("peer_suscription_id").(string)

	log.Printf("[INFO] Reading the status of VPC peering connection %s.", vpcPeeringID)
	vpcPeering, err := client.AzureReadVpcPeering(cdcID, vpcPeeringID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}
	err = MapAzureVPCPeeringToResource(d, vpcPeering)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}
	return nil

}

func MapAzureVPCPeeringToResource(d *schema.ResourceData, vpcPeering *AzureVPCPeering) error {
	if vpcPeering.ID == "" {
		return nil
	}
	d.SetId(vpcPeering.ID)
	d.Set("peer_suscription_id", vpcPeering.PeerSubscriptionId)
	d.Set("cdc_id", vpcPeering.ClusterDataCentre)
	d.Set("peer_resource_group", vpcPeering.PeerResourceGroup)
	d.Set("resource_group", vpcPeering.ResourceGroup)
	d.Set("peer_vpc_net", vpcPeering.PeerVNet)
	d.Set("suscription_id", vpcPeering.SubscriptionId)
	d.Set("v_net", vpcPeering.VNet)
	d.Set("peer_subnets", vpcPeering.PeerSubnets)

	log.Printf("[INFO] Fetched VPC peering %s info from the remote server.", vpcPeering.ID)
	return nil
}

func resourceAzureVpcPeeringUpdate(d *schema.ResourceData) error {
	return fmt.Errorf("[Error] The VPC peering connection doesn't support update")

}

func AzurecreateVpcPeeringRequest(d *schema.ResourceData) (CreateAzureVPCPeeringRequest, error) {
	result := CreateAzureVPCPeeringRequest{
		PeerVPCNetworkName: d.Get("peer_vpc_net").(string),
		PeerSubscriptionId: d.Get("peer_suscription_id").(string),
		PeerResourceGroup:  d.Get("peer_resource_group").(string),
	}
	if _, isSet := d.GetOk("peer_subnets"); isSet {
		result.PeerSubnets = d.Get("peer_subnets").(*schema.Set).List()
	} else {
		return result, fmt.Errorf("[Error] Error creating GCP VPC peering request - Please check the subnets atleast one subnet must be specified")
	}
	return result, nil
}
