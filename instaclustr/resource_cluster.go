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

var (
	validClusterStates = map[string]bool{
		"RUNNING":     true,
		"PROVISIONED": true,
	}
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
				ForceNew: true,
			},

			"node_size": {
				Type:     schema.TypeString,
				Required: true,
			},

			"data_centre": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"sla_tier": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NON_PRODUCTION",
				ForceNew: true,
			},

			"cluster_network": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10.224.0.0/12",
				ForceNew: true,
			},

			"private_network_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"pci_compliant_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
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
							ForceNew: true,
						},
						"account_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"custom_virtual_network_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"disk_encryption_key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},

			"rack_allocation": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number_of_racks": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"nodes_per_rack": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
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
					ForceNew: true,
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
							ForceNew: true,
						},
						"version": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"options": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_n_authz": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"client_encryption": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"use_private_broadcast_rpc_address": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"lucene_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"continuous_backup_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"number_partitions": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"auto_create_topics": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"delete_topics": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"password_authentication": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"target_kafka_cluster_id": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"vpc_type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"aws_access_key": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"aws_secret_key": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"s3_bucket_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"azure_storage_account_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"azure_storage_account_key": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"azure_storage_container_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"ssl_enabled_protocols": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"ssl_truststore_password": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"ssl_protocol": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"security_protocol": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"sasl_mechanism": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"sasl_jaas_config": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"bootstrap_servers": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"truststore": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"dedicated_zookeeper": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"zookeeper_node_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"zookeeper_node_count": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"master_nodes": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
									"replica_nodes": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"kafka_rest_proxy_user_password": {
				Type:     schema.TypeString,
				Sensitive: true,
				Optional: true,
			},
			"kafka_schema_registry_user_password": {
				Type:     schema.TypeString,
				Sensitive: true,
				Optional: true,
			},
			"minimum_required_cluster_state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",

				ValidateFunc: func(i interface{}, s string) (ws []string, errors []error) {
					state := i.(string)

					if len(state) != 0 && !validClusterStates[strings.ToUpper(state)] {
						errors = append(errors, fmt.Errorf("%s is not valid cluster state. Use RUNNING or PROVISIONED.", state))
					}

					return
				},
			},
		},
	}
}

var isKafkaCluster bool
var hasRestProxy bool
var hasSchemaRegistry bool

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

	var jsonStrCreate []byte
	jsonStrCreate, err = json.Marshal(createData)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	id, err := client.CreateCluster(jsonStrCreate)
	if err != nil {
		return formatCreateErrMsg(err)
	}
	d.SetId(id)
	d.Set("cluster_id", id)
	log.Printf("[INFO] Cluster %s has been created.", id)

	kafkaSchemaRegistryUserPassword := d.Get("kafka_schema_registry_user_password").(string)
	kafkaRestProxyUserPassword := d.Get("kafka_rest_proxy_user_password").(string)
	waitForClusterState := d.Get("minimum_required_cluster_state").(string)

	checkForBundleAvailability(bundles)

	if len(waitForClusterState) < 0 {
		return nil
	} else if (len(kafkaSchemaRegistryUserPassword) > 0 || len(kafkaRestProxyUserPassword) > 0) && waitForClusterState != "RUNNING" {
		return fmt.Errorf("[Error] Please specify the cluster to reach the RUNNING state before updating the kafka-schema-registry or kafka-rest-proxy user password with minimum_required_cluster_state property")
	} else {
		return waitForClusterStateAndDoUpdate(client, waitForClusterState, isKafkaCluster, hasRestProxy, hasSchemaRegistry, kafkaRestProxyUserPassword, kafkaSchemaRegistryUserPassword, d, id)
	}
}

func waitForClusterStateAndDoUpdate(client *APIClient,
	waitForClusterState string,
	isKafkaCluster bool,
	hasRestProxy bool,
	hasSchemaRegistry bool,
	kafkaRestProxyUserPassword string,
	kafkaSchemaRegistryUserPassword string,
	d *schema.ResourceData,
	id string) error {
	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {

		//reading cluster details
		cluster, err := client.ReadCluster(id)

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("[Error] Error retrieving cluster info: %s", err))
		}

		if cluster.ClusterStatus == waitForClusterState {

			if !isKafkaCluster && (len(kafkaSchemaRegistryUserPassword) > 0 || len(kafkaRestProxyUserPassword) > 0) {
				return resource.NonRetryableError(fmt.Errorf("[Error] Error updating the bundle user passwords, because it should be a KAFKA cluster in order to update the schema-registry or rest-proxy users"))
			}

			if isKafkaCluster && hasRestProxy && (len(kafkaRestProxyUserPassword) > 0) {
				updateKafkaRestProxyPassword(d, client)
			}

			if isKafkaCluster && hasSchemaRegistry && (len(kafkaSchemaRegistryUserPassword) > 0) {
				updateKafkaSchemaRegistryPassword(d, client)
			}
		} else {
			return resource.RetryableError(fmt.Errorf("[DEBUG] Cluster is in state %s, waiting for it to reach state %s", cluster.ClusterStatus, waitForClusterState))
		}
		return nil
	})
}

func updateKafkaRestProxyPassword(d *schema.ResourceData, client *APIClient) error {
	//updating the kafka rest proxy bundle user
	var err error
	err = client.UpdateKafkaRestProxyUser(d.Get("cluster_id").(string), createBundleUserUpdateRequest("ickafkarest", d.Get("kafka_rest_proxy_user_password").(string)))
	if err != nil {
		return fmt.Errorf("[Error] Error updating the kafka rest proxy bundle user password : %s", err)
	}

	return nil
}

func updateKafkaSchemaRegistryPassword(d *schema.ResourceData, client *APIClient) error {
	//updating the kafka schema registry bundle user
	var err error
	err = client.UpdateKafkaSchemaRegistryUser(d.Get("cluster_id").(string), createBundleUserUpdateRequest("ickafkaschema", d.Get("kafka_schema_registry_user_password").(string)))
	if err != nil {
		return fmt.Errorf("[Error] Error updating the kafka schema registry bundle user password : %s", err)
	}

	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// currently only cluster resize, kafka-schema-registry user password update and kafka-rest-proxy user password update are supported

	client := meta.(*Config).Client
	clusterID := d.Get("cluster_id").(string)

	clusterResize := d.HasChange("node_size")
	kafkaSchemaRegistryUserUpdate := d.HasChange("kafka_schema_registry_user_password")
	kafkaRestProxyUerUpdate := d.HasChange("kafka_rest_proxy_user_password")

	bundles, err := getBundles(d)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	checkForBundleAvailability(bundles)

	if isKafkaCluster && hasSchemaRegistry && kafkaSchemaRegistryUserUpdate {

		//updating the bundle user
		err = client.UpdateKafkaSchemaRegistryUser(clusterID, createBundleUserUpdateRequest("ickafkaschema", d.Get("kafka_schema_registry_user_password").(string)))
		if err != nil {
			return fmt.Errorf("[Error] Error updating the password for kafka schema registry user : %s", err)
		}
	}
	if isKafkaCluster && hasRestProxy && kafkaRestProxyUerUpdate {
		//updating the bundle user
		err = client.UpdateKafkaRestProxyUser(clusterID, createBundleUserUpdateRequest("ickafkarest", d.Get("kafka_rest_proxy_user_password").(string)))
		if err != nil {
			return fmt.Errorf("[Error] Error updating the password for kafka rest proxy user : %s", err)
		}
	}
	if clusterResize {
		doClusterResize(client, clusterID, d)
	}

	if !isKafkaCluster && (kafkaSchemaRegistryUserUpdate || kafkaRestProxyUerUpdate) {
		return fmt.Errorf("[Error] Error updating the bundle user passwords, because it should be a KAFKA cluster in order to update the schema-registry or rest-proxy users")
	}

	d.SetPartial("node_size")
	d.SetPartial("kafka_schema_registry_user_password")
	d.SetPartial("kafka_rest_proxy_user_password")
	return nil
}

func createBundleUserUpdateRequest(bundleUsername string, bundleUserPassword string) []byte {

	var err error
	//preparing the bundle user update request
	updateBundleUserData := UpdateBundleUserRequest{
		Username: bundleUsername,
		Password: bundleUserPassword,
	}
	var jsonStrUpdateBundleUser []byte
	jsonStrUpdateBundleUser, err = json.Marshal(updateBundleUserData)

	if err != nil {
		log.Printf("[ERROR] Error creating the bundle user update request : %s", err)
		return nil
	} else {
		return jsonStrUpdateBundleUser
	}
}

func checkForBundleAvailability(bundles []Bundle) {
	for i := 0; i < len(bundles); i++ {

		if bundles[i].Bundle == "KAFKA" {
			isKafkaCluster = true
		}
		if bundles[i].Bundle == "KAFKA_REST_PROXY" {
			hasRestProxy = true
		}
		if bundles[i].Bundle == "KAFKA_SCHEMA_REGISTRY" {
			hasSchemaRegistry = true
		}
	}
}

func doClusterResize(client *APIClient, clusterID string, d *schema.ResourceData) error {

	before, after := d.GetChange("node_size")
	regex := regexp.MustCompile(`resizeable-(small|large)`)
	oldNodeClass := regex.FindString(before.(string))
	newNodeClass := regex.FindString(after.(string))

	isNotResizable := (oldNodeClass == "")
	isNotSameSizeClass := (newNodeClass != oldNodeClass)
	if isNotResizable || isNotSameSizeClass {
		return fmt.Errorf("[Error] Cannot resize nodes from %s to %s", before, after)
	}

	cluster, err := client.ReadCluster(clusterID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}
	err = client.ResizeCluster(clusterID, cluster.DataCentres[0].ID, after.(string))
	if err != nil {
		return fmt.Errorf("[Error] Error resizing cluster %s with error %s", clusterID, err)
	}
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

	nodeSize := ""
	/* 
	*  Ideally, we would like this information to be coming directly from the API cluster status.
	*  Hence, this is a slightly hacky way of ignoring zookeeper node sizes (Kafka bundle specific).
	*/
	for _, node := range(cluster.DataCentres[0].Nodes) {
		nodeSize = node.Size
		if (!strings.HasPrefix(nodeSize, "zk-")) {
			break
		}
	}
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

	toCheck := [3]string{"cluster_provider","rack_allocation","bundle"}
	for _, changing := range toCheck {
		if d.HasChange(changing) {
			_, after := d.GetChange(changing)
			d.Set(changing, after)
		}
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
		err := mapstructure.WeakDecode(inBundle.(map[string]interface{}), &bundle)
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
