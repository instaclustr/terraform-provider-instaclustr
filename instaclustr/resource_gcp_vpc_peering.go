package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGCPVpcPeering() *schema.Resource {
	return &schema.Resource{
		Create: GCPresourceVpcPeeringCreate,
		Read:   GCPresourceVpcPeeringRead,
		Update: resourceVpcPeeringUpdate,
		Delete: GCPresourceVpcPeeringDelete,

		Importer: &schema.ResourceImporter{
			State: GCPresourceVpcPeeringStateImport,
		},

		Schema: map[string]*schema.Schema{

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
		},
	}
}

func GCPresourceVpcPeeringCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating GCP VPC peering Connection request.")
	client := meta.(*Config).Client

	const ClusterReadInterval = 5
	const WaitForClusterTimeout = 60
	var cdcID string
	var latestStatus string
	timePassed := 0
	for {
		cluster, err := client.ReadCluster(d.Get("cluster_id").(string))
		if err != nil {
			return fmt.Errorf("[Error] Error in retrieving the cluster info: %s", err)
		}
		latestStatus = cluster.ClusterStatus
		if cluster.DataCentres[0].CdcStatus == "PROVISIONED!" || cluster.ClusterStatus == "RUNNING!" {
			cdcID = cluster.DataCentres[0].ID
			break
		}
		if timePassed > WaitForClusterTimeout {
			return fmt.Errorf("[Error] Timed out, waiting for cluster status to be 'PROVISIONED' or 'RUNNING'. Current cluster status is '%s'", latestStatus)
		}
		time.Sleep(ClusterReadInterval * time.Second)
		timePassed += ClusterReadInterval
	}

	createData, err := GCPcreateVpcPeeringRequest(d)
	if err != nil {
		return fmt.Errorf("[Error] Error creating GCP VPC peering request: %s", err)
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(createData)
	if err != nil {
		return fmt.Errorf("[Error] Error creating GCP VPC peering request: %s", err)
	}

	id, err := client.CreateVpcPeering(cdcID, jsonStr)
	if err != nil {
		return fmt.Errorf("[Error] Error in creating the cluster: %s", err)
	}
	d.SetId(id)
	d.Set("vpc_peering_id", id)
	d.Set("cdc_id", cdcID)

	vpcPeering, err := client.GCPReadVpcPeering(cdcID, id)
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

	d.SetId(vpcPeering.ID)
	d.Set("vpc_peering_id", vpcPeering.ID)
	d.Set("cdc_id", vpcPeering.ClusterDataCentre)
	d.Set("peer_vpc_network_name", vpcPeering.PeerVPCNetworkName)

	d.Set("peer_project_id", vpcPeering.PeerProjectID)

	log.Printf("[INFO] Fetched VPC peering %s info from the remote server.", vpcPeering.ID)
	return nil

}

func resourceGCPVpcPeeringUpdate(d *schema.ResourceData) error {
	return fmt.Errorf("[Error] The VPC peering connection doesn't support update")

}

func GCPresourceVpcPeeringDelete(d *schema.ResourceData, meta interface{}) error {
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

func GCPresourceVpcPeeringStateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "&")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("Unexpected format of ID (%q), expected <CLUSTER-ID>&<VPC-PEERING-ID>", d.Id())
	}

	d.Set("cluster_id", idParts[0])
	d.Set("vpc_peering_id", idParts[1])
	return []*schema.ResourceData{d}, nil
}

func GCPcreateVpcPeeringRequest(d *schema.ResourceData) (CreateGCPVPCPeeringRequest, error) {
	result := CreateGCPVPCPeeringRequest{
		Name:               d.Get("name").(string),
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
