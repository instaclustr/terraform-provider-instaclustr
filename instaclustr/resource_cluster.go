package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"cluster_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"node_size": {
				Type:     schema.TypeString,
				Required: true,
			},

			"data_centre": {
				Type:     schema.TypeString,
				Required: true,
			},

			"sla_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NON_PRODUCTION",
			},

			"cluster_network": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10.224.0.0/12",
			},

			"private_network_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"pci_compliant_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"public_contact_point": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"private_contact_point": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"cluster_provider": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"account_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tags": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"disk_encryption_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"rack_allocation": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number_of_racks": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"nodes_per_rack": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},

			"bundles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: schema.TypeString,
				},
				Removed: "Please change bundles argument -> bundle blocks (example under example/main.tf), and to avoid causing an update to the existing tfstate - replace all keys named 'bundles' with 'bundle' in resources with the provider 'provider.instaclustr'",
			},

			"bundle": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bundle": {
							Type:     schema.TypeString,
							Required: true,
						},
						"version": {
							Type:     schema.TypeString,
							Required: true,
						},
						"options": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_n_authz": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"client_encryption": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"use_private_broadcast_rpc_address": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"lucene_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"continuous_backup_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"number_partitions": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"auto_create_topics": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"delete_topics": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"password_authentication": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceClusterCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Creating cluster.")
	client := meta.(*Config).Client

	bundles, err := getBundles(d)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	var clusterProvider ClusterProvider
	err = mapstructure.Decode(d.Get("cluster_provider").(map[string]interface{}), &clusterProvider)
	if err != nil {
		return err
	}

	var rackAllocation RackAllocation
	err = mapstructure.Decode(d.Get("rack_allocation").(map[string]interface{}), &rackAllocation)
	if err != nil {
		return err
	}

	createData := CreateRequest{
		ClusterName:           d.Get("cluster_name").(string),
		Bundles:               bundles,
		Provider:              clusterProvider,
		SlaTier:               d.Get("sla_tier").(string),
		NodeSize:              d.Get("node_size").(string),
		DataCentre:            d.Get("data_centre").(string),
		ClusterNetwork:        d.Get("cluster_network").(string),
		PrivateNetworkCluster: fmt.Sprintf("%v", d.Get("private_network_cluster")),
		PCICompliantCluster:   fmt.Sprintf("%v", d.Get("pci_compliant_cluster"))
		RackAllocation:        rackAllocation,
	}

	var jsonStr []byte
	jsonStr, err = json.Marshal(createData)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	id, err := client.CreateCluster(jsonStr)
	if err != nil {
		return formatCreateErrMsg(err)
	}
	d.SetId(id)
	d.Set("cluster_id", id)
	log.Printf("[INFO] Cluster %s has been created.", id)
	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// currently only cluster resize is supported
	if !d.HasChange("node_size") {
		return fmt.Errorf("[Error] The cluster doesn't support update")
	}

	before, after := d.GetChange("node_size")
	regex := regexp.MustCompile(`resizeable-(small|large)`)
	oldNodeClass := regex.FindString(before.(string))
	newNodeClass := regex.FindString(after.(string))

	isNotResizable := (oldNodeClass == "")
	isNotSameSizeClass := (newNodeClass != oldNodeClass)
	if isNotResizable || isNotSameSizeClass {
		return fmt.Errorf("[Error] Cannot resize nodes from %s to %s", before, after)
	}

	client := meta.(*Config).Client
	clusterID := d.Get("cluster_id").(string)
	cluster, err := client.ReadCluster(clusterID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}
	err = client.ResizeCluster(clusterID, cluster.DataCentres[0].ID, after.(string))
	if err != nil {
		return fmt.Errorf("[Error] Error resizing cluster %s with error %s", clusterID, err)
	}

	d.SetPartial("node_size")
	return nil
}

func resourceClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)
	log.Printf("[INFO] Reading status of cluster %s.", id)
	cluster, err := client.ReadCluster(id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}
	d.SetId(cluster.ID)
	d.Set("cluster_id", cluster.ID)
	d.Set("cluster_name", cluster.ClusterName)

	nodeSize := cluster.DataCentres[0].Nodes[0].Size
	if len(cluster.DataCentres[0].ResizeTargetNodeSize) > 0 {
		nodeSize = cluster.DataCentres[0].ResizeTargetNodeSize
	}
	d.Set("node_size", nodeSize)

	d.Set("data_centre", cluster.DataCentres[0].Name)
	d.Set("sla_tier", strings.ToUpper(cluster.SlaTier))
	d.Set("cluster_network", cluster.DataCentres[0].CdcNetwork)
	d.Set("private_network_cluster", cluster.DataCentres[0].PrivateIPOnly)
	d.Set("pci_compliant_cluster", cluster.PciCompliance)

	if len(cluster.DataCentres[0].Nodes[0].PublicAddress) != 0 {
		err = d.Set("public_contact_point", cluster.DataCentres[0].Nodes[0].PublicAddress)

	}

	if len(cluster.DataCentres[0].Nodes[0].PrivateAddress) != 0 {
		err = d.Set("private_contact_point", cluster.DataCentres[0].Nodes[0].PrivateAddress)
	}

	var before interface{}
	before, _ = d.GetChange("cluster_provider")
	d.Set("cluster_provider", before)
	before, _ = d.GetChange("rack_allocation")
	d.Set("rack_allocation", before)
	before, _ = d.GetChange("bundle")
	d.Set("bundle", before)
	log.Printf("[INFO] Fetched cluster %s info from the remote server.", cluster.ID)
	return nil
}

func resourceClusterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Get("cluster_id").(string)
	log.Printf("[INFO] Deleting cluster %s.", id)
	err := client.DeleteCluster(id)
	if err != nil {
		return fmt.Errorf("[Error] Error deleting cluster: %s", err)
	}
	d.SetId("")
	d.Set("cluster_id", "")
	log.Printf("[INFO] Cluster %s has been marked for deletion.", id)
	return nil
}

func getBundles(d *schema.ResourceData) ([]Bundle, error) {
	bundles := make([]Bundle, 0)
	for _, inBundle := range d.Get("bundle").([]interface{}) {
		var bundle Bundle
		err := mapstructure.Decode(inBundle.(map[string]interface{}), &bundle)
		if err != nil {
			return nil, err
		}
		bundles = append(bundles, bundle)
	}
	return bundles, nil
}

func formatCreateErrMsg(err error) error {
	return fmt.Errorf("[Error] Error creating cluster: %s", err)
}
