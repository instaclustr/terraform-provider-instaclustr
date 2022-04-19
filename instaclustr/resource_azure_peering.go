package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAzureVpcPeering() *schema.Resource {
	return &schema.Resource{
		Create: azureResourceVpcPeeringCreate,
		Read:   azureResourceVpcPeeringRead,
		Update: resourceVpcPeeringUpdate,
		Delete: resourceAzureVpcPeeringDelete,

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

			"peer_subscription_id": {
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

			"subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"peer_subnets": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},

			"virtual_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func azureResourceVpcPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	cdcID, err := VpcPeeringCreate(d, meta)
	if err != nil {
		return fmt.Errorf("[Error] Error creating Azure VPC peering request object: %s", err)
	}
	var createData CreateAzureVPCPeeringRequest
	createData, err = azureCreateVpcPeeringRequest(d)
	if err != nil {
		return fmt.Errorf("[Error] Error creating Azure VPC peering request: %s", err)
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating Azure VPC peering request: %s", err)
	}
	var id string
	id, err = client.CreateVpcPeering(cdcID, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error in creating peering request to the cluster: %s", err)
	}
	d.SetId(id)
	d.Set("peer_subscription_id", id)
	d.Set("cdc_id", cdcID)
	var vpcPeering *AzureVPCPeering
	vpcPeering, err = client.AzureReadVpcPeering(cdcID, id)
	if err != nil {
		return fmt.Errorf("[Error] Error in reading Azure VPC peering connection: %s", err)
	}

	d.Set("peer_subscription_id", vpcPeering.PeerSubscriptionId)

	log.Printf("[INFO] VPC peering request %s has been created.", id)
	return nil
}

func azureResourceVpcPeeringRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
	}
	cdcID := cluster.DataCentres[0].ID
	vpcPeeringID := d.Id()

	log.Printf("[INFO] Reading the status of VPC peering connection %s.", vpcPeeringID)
	vpcPeering, err := client.AzureReadVpcPeering(cdcID, vpcPeeringID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}
	err = mapAzureVPCPeeringToResource(d, vpcPeering)
	if err != nil {
		return fmt.Errorf("[Error] Error reading VPC peering connection: %s", err)
	}
	return nil
}

func mapAzureVPCPeeringToResource(d *schema.ResourceData, vpcPeering *AzureVPCPeering) error {
	if vpcPeering.ID == "" {
		return fmt.Errorf("[Error] Error creating Azure VPC peering request - Please check the cluster ID")
	}
	d.SetId(vpcPeering.ID)
	d.Set("peer_subscription_id", vpcPeering.PeerSubscriptionId)
	d.Set("cdc_id", vpcPeering.ClusterDataCentre)
	d.Set("peer_resource_group", vpcPeering.PeerResourceGroup)
	d.Set("resource_group", vpcPeering.ResourceGroup)
	d.Set("peer_vpc_net", vpcPeering.PeerVNet)
	d.Set("subscription_id", vpcPeering.SubscriptionId)
	d.Set("virtual_network_id", vpcPeering.VirtualNetworkID)
	d.Set("peer_subnets", vpcPeering.PeerSubnets)

	log.Printf("[INFO] Fetched VPC peering %s info from the remote server.", vpcPeering.ID)
	return nil
}

func resourceAzureVpcPeeringDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
	if err != nil {
		return fmt.Errorf("[Error] Error retrieving cluster info: %s", err)
	}
	cdcID := cluster.DataCentres[0].ID
	vpcPeeringID := d.Id()
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

func azureCreateVpcPeeringRequest(d *schema.ResourceData) (CreateAzureVPCPeeringRequest, error) {
	result := CreateAzureVPCPeeringRequest{
		PeerVPCNetworkName: d.Get("peer_vpc_net").(string),
		PeerSubscriptionId: d.Get("peer_subscription_id").(string),
		PeerResourceGroup:  d.Get("peer_resource_group").(string),
		PeerSubnets:        d.Get("peer_subnets").(*schema.Set).List(),
	}
	return result, nil
}
