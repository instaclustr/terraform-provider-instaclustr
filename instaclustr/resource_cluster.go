package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
		},

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

			"data_centre_id": {
				Type:     schema.TypeString,
				Computed: true,
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

			"public_contact_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"private_contact_point": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"private_contact_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"cluster_certificate_download": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"username": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},

			"instaclustr_user_password": {
				Type:      schema.TypeString,
				Computed:  true,
				Optional:  true,
				Sensitive: true,
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
						"custom_virtual_network_id": {
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

			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"rack_allocation": {
				Type:     schema.TypeMap,
				Optional: true,
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
									"target_kafka_cluster_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"vpc_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"aws_access_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"aws_secret_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"s3_bucket_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_storage_account_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_storage_account_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"azure_storage_container_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssl_enabled_protocols": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssl_truststore_password": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ssl_protocol": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"security_protocol": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sasl_mechanism": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sasl_jaas_config": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"bootstrap_servers": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"truststore": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"dedicated_zookeeper": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"zookeeper_node_size": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"zookeeper_node_count": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"master_nodes": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"replica_nodes": {
										Type:     schema.TypeInt,
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

	clusterProvider.Tags = d.Get("tags").(map[string]interface{})

	createData := CreateRequest{
		ClusterName:           d.Get("cluster_name").(string),
		Bundles:               bundles,
		Provider:              clusterProvider,
		SlaTier:               d.Get("sla_tier").(string),
		NodeSize:              d.Get("node_size").(string),
		DataCentre:            d.Get("data_centre").(string),
		ClusterNetwork:        d.Get("cluster_network").(string),
		PrivateNetworkCluster: fmt.Sprintf("%v", d.Get("private_network_cluster")),
		PCICompliantCluster:   fmt.Sprintf("%v", d.Get("pci_compliant_cluster")),
	}

	// Some Bundles do not use Rack Allocation so add that separately if needed. (Redis for example)
	if checkIfBundleRequiresRackAllocation(bundles) {
		var rackAllocation RackAllocation
		err = mapstructure.Decode(d.Get("rack_allocation").(map[string]interface{}), &rackAllocation)
		if err != nil {
			return err
		}

		createData.RackAllocation = &rackAllocation
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


	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		cluster, err := client.ReadCluster(id)

		if err != nil {
				return resource.NonRetryableError(fmt.Errorf("[Error] Error retrieving cluster info: %s", err))
			}

		if cluster.ClusterStatus == "RUNNING" {
				return resource.NonRetryableError(resourceClusterRead(d, meta))
			}

		return resource.RetryableError(fmt.Errorf("[Error] Cluster is in state %s", cluster.ClusterStatus))
	})

	return resourceClusterRead(d, meta)

}

func resourceClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// currently only cluster resize is supported
	if d.HasChange("node_size") {
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
	}
	return resourceClusterRead(d, meta)
}

func resourceClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client
	id := d.Id()
	log.Printf("[INFO] Reading status of cluster %s.", id)
	cluster, err := client.ReadCluster(id)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}
	d.SetId(cluster.ID)
	d.Set("cluster_id", cluster.ID)
	d.Set("cluster_name", cluster.ClusterName)

	nodeSize := ""
	/*
	*  Ideally, we would like this information to be coming directly from the API cluster status.
	*  Hence, this is a slightly hacky way of ignoring zookeeper node sizes (Kafka bundle specific).
	 */
	for _, node := range cluster.DataCentres[0].Nodes {
		nodeSize = node.Size
		if !strings.HasPrefix(nodeSize, "zk-") {
			break
		}
	}
	if len(cluster.DataCentres[0].ResizeTargetNodeSize) > 0 {
		nodeSize = cluster.DataCentres[0].ResizeTargetNodeSize
	}
	d.Set("node_size", nodeSize)

	d.Set("data_centre", cluster.DataCentres[0].Name)
	d.Set("data_centre_id", cluster.DataCentres[0].ID)
	d.Set("sla_tier", strings.ToUpper(cluster.SlaTier))
	d.Set("cluster_network", cluster.DataCentres[0].CdcNetwork)
	d.Set("private_network_cluster", cluster.DataCentres[0].PrivateIPOnly)
	d.Set("pci_compliant_cluster", cluster.PciCompliance)
	d.Set("cluster_certificate_download", cluster.ClusterCertificateDownload)
	d.Set("username", cluster.Username)

	if len(cluster.InstaclustrUserPassword) > 0 {
		d.Set("instaclustr_user_password", cluster.InstaclustrUserPassword)
	}

	if len(cluster.DataCentres[0].Nodes[0].PublicAddress) != 0 {
		err = d.Set("public_contact_point", cluster.DataCentres[0].Nodes[0].PublicAddress)
		var nodes []string
		for _, node := range cluster.DataCentres[0].Nodes {
			nodes = append(nodes, node.PublicAddress)
		}
		err = d.Set("public_contact_addresses", nodes)
	}

	if len(cluster.DataCentres[0].Nodes[0].PrivateAddress) != 0 {
		err = d.Set("private_contact_point", cluster.DataCentres[0].Nodes[0].PrivateAddress)
		var nodes []string
		for _, node := range cluster.DataCentres[0].Nodes {
			nodes = append(nodes, node.PrivateAddress)
		}
		err = d.Set("private_contact_addresses", nodes)
	}

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

func checkIfBundleRequiresRackAllocation(bundles []Bundle) bool {
	var noRackAllocationBundles = []string{
		"REDIS",
	}

	for i := 0; i < len(bundles); i++ {
		for j := 0; j < len(noRackAllocationBundles); j++ {
			if strings.ToLower(bundles[i].Bundle) == strings.ToLower(noRackAllocationBundles[j]) {
				return false
			}
		}
	}

	return true
}
