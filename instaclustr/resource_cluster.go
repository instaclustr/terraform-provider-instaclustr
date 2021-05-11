package instaclustr

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
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

		Importer: &schema.ResourceImporter{
			State: resourceClusterStateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
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
				Optional: true,
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
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"private_contact_point": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
					Type:     schema.TypeMap,
					Elem:     schema.TypeString,
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
									"dedicated_master_nodes": {
										Type:     schema.TypeBool,
										Optional: false,
										ForceNew: true,
									},
									"master_node_size": {
										Type:     schema.TypeString,
										Optional: false,
										ForceNew: true,
									},
									"data_node_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"kibana_node_size": {
										Type:     schema.TypeString,
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
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"kafka_schema_registry_user_password": {
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			"wait_for_state": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func getNodeSize(d *schema.ResourceData, bundles []Bundle) (string, error) {
	for i, bundle := range bundles {
		if bundle.Bundle == "ELASTICSEARCH" {
			if len(bundle.Options.MasterNodeSize) == 0 {
				return "", fmt.Errorf("[ERROR] 'master_node_size' is required in the bundle option.")
			}
			dedicatedMaster := bundle.Options.DedicatedMasterNodes != nil && *bundle.Options.DedicatedMasterNodes
			if dedicatedMaster {
				if len(bundle.Options.DataNodeSize) == 0 {
					return "", fmt.Errorf("[ERROR] Elasticsearch dedicated master is enabled, 'data_node_size' is required in the bundle option.")
				}
				return bundle.Options.DataNodeSize, nil
			} else {
				if len(bundle.Options.DataNodeSize) != 0 && bundle.Options.DataNodeSize != bundle.Options.MasterNodeSize {
					return "", fmt.Errorf("[ERROR] When 'dedicated_master_nodes' is not true , data_node_size can be either null or equal to master_node_size.")
				}
				size := bundle.Options.MasterNodeSize
				bundles[i].Options.MasterNodeSize = ""
				return size, nil
			}
		}
	}
	if size, ok := d.GetOk("node_size"); !ok || len(size.(string)) == 0 {
		return "", fmt.Errorf("[ERROR] node_size must be set.")
	} else {
		return size.(string), nil
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

	size, err := getNodeSize(d, bundles)
	if err != nil {
		return err
	}
	createData := CreateRequest{
		ClusterName:           d.Get("cluster_name").(string),
		Bundles:               bundles,
		Provider:              clusterProvider,
		SlaTier:               d.Get("sla_tier").(string),
		NodeSize:              size,
		DataCentre:            d.Get("data_centre").(string),
		ClusterNetwork:        d.Get("cluster_network").(string),
		PrivateNetworkCluster: fmt.Sprintf("%v", d.Get("private_network_cluster")),
		PCICompliantCluster:   fmt.Sprintf("%v", d.Get("pci_compliant_cluster")),
	}

	kafkaSchemaRegistryUserPassword := d.Get("kafka_schema_registry_user_password").(string)
	kafkaRestProxyUserPassword := d.Get("kafka_rest_proxy_user_password").(string)
	waitForClusterState := d.Get("wait_for_state").(string)

	bundleConfig := getBundleConfig(bundles)

	if (len(kafkaSchemaRegistryUserPassword) > 0 || len(kafkaRestProxyUserPassword) > 0) && waitForClusterState != "RUNNING" {
		return fmt.Errorf("[Error] wait_for_state must be set as RUNNING when providing the kafka-schema-registry or kafka-rest-proxy user password")
	}

	if !bundleConfig.IsKafkaCluster && (len(kafkaSchemaRegistryUserPassword) > 0 || len(kafkaRestProxyUserPassword) > 0) {
		return fmt.Errorf("[Error] kafka-schema-registry or kafka-rest-proxy user passwords may only be provided for Kafka clusters")
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

	if len(waitForClusterState) == 0 {
		return nil
	}

	return waitForClusterStateAndDoUpdate(client, waitForClusterState, bundleConfig, kafkaRestProxyUserPassword, kafkaSchemaRegistryUserPassword, d, id, meta)
}

func waitForClusterStateAndDoUpdate(client *APIClient,
	waitForClusterState string,
	bundleConfig BundleConfig,
	kafkaRestProxyUserPassword string,
	kafkaSchemaRegistryUserPassword string,
	d *schema.ResourceData,
	id string,
	meta interface{}) error {
	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		fmt.Printf("waiting for cluster to reach %s\n", waitForClusterState)
		//reading cluster details
		cluster, err := client.ReadCluster(id)
		resourceClusterRead(d, meta)

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("[Error] Error retrieving cluster info: %s", err))
		}

		if cluster.ClusterStatus != waitForClusterState {
			return resource.RetryableError(fmt.Errorf("[DEBUG] Cluster is in state %s, waiting for it to reach state %s", cluster.ClusterStatus, waitForClusterState))
		}

		if bundleConfig.IsKafkaCluster && bundleConfig.HasRestProxy && (len(kafkaRestProxyUserPassword) > 0) {
			err = client.UpdateBundleUser(d.Get("cluster_id").(string), "kafka_rest_proxy", createBundleUserUpdateRequest("ickafkarest", d.Get("kafka_rest_proxy_user_password").(string)))
			if err != nil {
				return resource.RetryableError(fmt.Errorf("[DEBUG] Error updating the kafka rest proxy bundle user password : %s", err))
			}
		}

		if bundleConfig.IsKafkaCluster && bundleConfig.HasSchemaRegistry && (len(kafkaSchemaRegistryUserPassword) > 0) {
			err = client.UpdateBundleUser(d.Get("cluster_id").(string), "kafka_schema_registry", createBundleUserUpdateRequest("ickafkaschema", d.Get("kafka_schema_registry_user_password").(string)))
			if err != nil {
				return resource.RetryableError(fmt.Errorf("[DEBUG] Error updating the kafka schema registry bundle user password : %s", err))
			}
		}

		return nil
	})
}

func waitForCdcResizeToFinish(client *APIClient,
	opId string,
	clusterId string,
	cdcId string,
	d *schema.ResourceData) error {
	return resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
		res, err := client.GetCDCResizeDetail(opId, clusterId, cdcId)
		if err != nil {
			return resource.RetryableError(err)
		}
		if res.CompletedStatus == "FAILED" {
			return resource.NonRetryableError(fmt.Errorf("[Error] CDC Resize operation %s failed", opId))
		}
		if res.Completed == nil {
			return resource.RetryableError(fmt.Errorf("[DEBUG] CDC resize %s is in progress...", opId))
		}
		if res.CompletedStatus == "SUCCESS" {
			return nil
		} else {
			return resource.NonRetryableError(fmt.Errorf("[ERROR] unexpected cdc resize finish status %s", res.CompletedStatus))
		}
	})
}

func resourceClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// currently only cluster resize, kafka-schema-registry user password update and kafka-rest-proxy user password update are supported

	client := meta.(*Config).Client
	clusterID := d.Get("cluster_id").(string)

	kafkaSchemaRegistryUserUpdate := d.HasChange("kafka_schema_registry_user_password")
	kafkaRestProxyUserUpdate := d.HasChange("kafka_rest_proxy_user_password")

	bundles, err := getBundles(d)
	if err != nil {
		return formatCreateErrMsg(err)
	}

	bundleConfig := getBundleConfig(bundles)

	if bundleConfig.IsKafkaCluster && bundleConfig.HasSchemaRegistry && kafkaSchemaRegistryUserUpdate {
		//updating the bundle user
		err = client.UpdateBundleUser(clusterID, "kafka_schema_registry", createBundleUserUpdateRequest("ickafkaschema", d.Get("kafka_schema_registry_user_password").(string)))
		if err != nil {
			return fmt.Errorf("[Error] Error updating the password for kafka schema registry user : %s", err)
		}
	}

	if bundleConfig.IsKafkaCluster && bundleConfig.HasRestProxy && kafkaRestProxyUserUpdate {
		//updating the bundle user
		err = client.UpdateBundleUser(clusterID, "kafka_rest_proxy", createBundleUserUpdateRequest("ickafkarest", d.Get("kafka_rest_proxy_user_password").(string)))
		if err != nil {
			return fmt.Errorf("[Error] Error updating the password for kafka rest proxy user : %s", err)
		}
	}

	err = doClusterResize(client, clusterID, d, bundles)
	if err != nil {
		return fmt.Errorf("[Error] Error resizing the cluster : %s", err)
	}

	if !bundleConfig.IsKafkaCluster && (kafkaSchemaRegistryUserUpdate || kafkaRestProxyUserUpdate) {
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
	}
	return jsonStrUpdateBundleUser
}

func getBundleConfig(bundles []Bundle) BundleConfig {
	configs := BundleConfig{
		IsKafkaCluster:    false,
		HasRestProxy:      false,
		HasSchemaRegistry: false,
	}

	for i := 0; i < len(bundles); i++ {

		if bundles[i].Bundle == "KAFKA" {
			configs.IsKafkaCluster = true
		}
		if bundles[i].Bundle == "KAFKA_REST_PROXY" {
			configs.HasRestProxy = true
		}
		if bundles[i].Bundle == "KAFKA_SCHEMA_REGISTRY" {
			configs.HasSchemaRegistry = true
		}
	}
	return configs
}

func appendIfMissing(slice []string, toAppend string) []string {
	for _, element := range slice {
		if element == toAppend {
			return slice
		}
	}
	return append(slice, toAppend)
}

func getBundleIndex(bundleType string, bundles []Bundle) (int, error) {
	for i := range bundles {
		if bundles[i].Bundle == bundleType {
			return i, nil
		}
	}
	return -1, fmt.Errorf("can't find bundle %s", bundleType)
}

func hasElasticsearchSizeChanges(bundleIndex int, d *schema.ResourceData) bool {
	return len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "master_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "kibana_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "data_node_size"))) > 0
}

func hasKafkaSizeChanges(bundleIndex int, d *schema.ResourceData) bool {
	return len(getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "zookeeper_node_size"))) > 0 ||
		len(getNewSizeOrEmpty(d, "node_size")) > 0
}

func hasCassandraSizeChanges(d *schema.ResourceData) bool {
	return len(getNewSizeOrEmpty(d, "node_size")) > 0
}

func doClusterResize(client *APIClient, clusterID string, d *schema.ResourceData, bundles []Bundle) error {
	cluster, err := client.ReadCluster(clusterID)
	if err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}
	bundleIndex, err := getBundleIndex(cluster.BundleType, bundles)
	if err != nil {
		return err
	}
	var res *ClusterDataCenterResizeResponse
	var cdcId *string
	switch cluster.BundleType {
	case "APACHE_CASSANDRA":
		if hasCassandraSizeChanges(d) {
			cdcId, res, err = doLegacyCassandraClusterResize(client, cluster, d)
		} else {
			return nil
		}
		break
	case "ELASTICSEARCH":
		if hasElasticsearchSizeChanges(bundleIndex, d) {
			cdcId, res, err = doElasticsearchClusterResize(client, cluster, d, bundleIndex)
		} else {
			return nil
		}
		break
	case "KAFKA":
		if hasKafkaSizeChanges(bundleIndex, d) {
			cdcId, res, err = doKafkaClusterResize(client, cluster, d, bundleIndex)
		} else {
			return nil
		}
		break
	default:
		return fmt.Errorf("CDC resize does not support: %s", cluster.BundleType)
	}
	if err != nil {
		return err
	}
	return waitForCdcResizeToFinish(client, res.OperationId, clusterID, *cdcId, d)
}

func getNewSizeOrEmpty(d *schema.ResourceData, key string) string {
	if !d.HasChange(key) {
		return ""
	}
	_, after := d.GetChange(key)
	return after.(string)
}

func isElasticsearchSizeAllChange(kibanaSize, masterSize, dataSize string, kibana, dataNodes bool) (string, bool) {
	if len(masterSize) == 0 {
		return "", false
	}
	if kibana && (len(kibanaSize) == 0 || kibanaSize != masterSize) {
		return "", false
	}
	if dataNodes && (len(dataSize) == 0 || dataSize != masterSize) {
		return "", false
	}
	return masterSize, true
}

func getSingleChangedElasticsearchSizeAndPurpose(kibanaSize, masterSize, dataSize string, kibana, dataNodes bool) (string, NodePurpose, error) {
	changedCount := 0
	var nodePurpose NodePurpose
	var nodeSize string
	if len(masterSize) > 0 {
		changedCount += 1
		nodePurpose = ELASTICSEARCH_MASTER
		nodeSize = masterSize
	}
	if len(dataSize) > 0 {
		if !dataNodes {
			return "", "", fmt.Errorf("[ERROR] This cluster does not have data only nodes, so data_node_sise is not used for this cluster. Please use master_node_size to change the size instead")
		}
		changedCount += 1
		nodePurpose = ELASTICSEARCH_DATA_AND_INGEST
		nodeSize = dataSize
	}
	if len(kibanaSize) > 0 {
		if !kibana {
			return "", "", fmt.Errorf("[ERROR] This cluster didn't enable Kibana, kibana_node_sise is not used for this cluster. Please use master_node_size to change the size instead")
		}
		changedCount += 1
		nodePurpose = ELASTICSEARCH_KIBANA
		nodeSize = kibanaSize
	}
	if changedCount > 1 {
		return "", "", fmt.Errorf("[ERROR] Please change one size at a time or change all to same size at onece")
	}
	return nodeSize, nodePurpose, nil
}

func getBundleOptionKey(bundleIndex int, option string) string {
	return fmt.Sprintf("bundle.%d.options.%s", bundleIndex, option)
}

func doElasticsearchClusterResize(client *APIClient, cluster *Cluster, d *schema.ResourceData, bundleIndex int) (*string, *ClusterDataCenterResizeResponse, error) {
	kibana := len(cluster.BundleOption.KibanaNodeSize) > 0
	dataNodes := len(cluster.BundleOption.DataNodeSize) > 0
	var nodePurpose *NodePurpose
	var nodeSize string
	masterNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "master_node_size"))
	kibanaNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "kibana_node_size"))
	dataNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "data_node_size"))

	if newSize, isAllChange := isElasticsearchSizeAllChange(kibanaNewSize, masterNewSize, dataNewSize, kibana, dataNodes); isAllChange {
		nodeSize = newSize
		nodePurpose = nil
	} else {
		newSize, purpose, err := getSingleChangedElasticsearchSizeAndPurpose(kibanaNewSize, masterNewSize, dataNewSize, kibana, dataNodes)
		if err != nil {
			return nil, nil, err
		}
		nodeSize = newSize
		nodePurpose = &purpose
	}
	log.Printf("[INFO] Resizing Elasticsearch cluster. nodePurpose: %s, newSize: %s", nodePurpose, nodeSize)
	res, err := client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, nodeSize, nodePurpose)
	if err != nil {
		return nil, nil, fmt.Errorf("[Error] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return &cluster.DataCentres[0].ID, res, nil
}

func isKafkaSizeAllChange(brokerSize, zookeeperSize string, dedicatedZookeeper bool) (string, bool) {
	if len(brokerSize) == 0 {
		return "", false
	}
	if dedicatedZookeeper && (len(zookeeperSize) == 0 || zookeeperSize != brokerSize) {
		return "", false
	}
	return brokerSize, true
}

func getSingleChangedKafkaSizeAndPurpose(brokerSize, zookeeperSize string, dedicatedZookeeper bool) (string, NodePurpose, error) {
	changedCount := 0
	var nodePurpose NodePurpose
	var nodeSize string
	if len(brokerSize) > 0 {
		changedCount += 1
		nodePurpose = KAFKA_BROKER
		nodeSize = brokerSize
	}
	if len(zookeeperSize) > 0 {
		if !dedicatedZookeeper {
			return "", "", fmt.Errorf("[ERROR] This cluster didn't enable Dedicated Zookeeper, zookeeper_node_size is not used for this cluster. Please use master_node_size to change the size instead")
		}
		nodePurpose = KAFKA_DEDICATED_ZOOKEEPER
		nodeSize = zookeeperSize
		changedCount += 1
	}
	if changedCount > 1 {
		return "", "", fmt.Errorf("[ERROR] Please change one size at a time or change all to same size at onece")
	}
	return nodeSize, nodePurpose, nil
}

func doKafkaClusterResize(client *APIClient, cluster *Cluster, d *schema.ResourceData, bundleIndex int) (*string, *ClusterDataCenterResizeResponse, error) {
	dedicatedZookeeper := cluster.BundleOption.DedicatedZookeeper != nil && *cluster.BundleOption.DedicatedZookeeper

	var nodePurpose *NodePurpose
	var nodeSize string
	zookeeperNewSize := getNewSizeOrEmpty(d, getBundleOptionKey(bundleIndex, "zookeeper_node_size"))
	brokerNewSize := getNewSizeOrEmpty(d, "node_size")

	if newSize, isAllChange := isKafkaSizeAllChange(brokerNewSize, zookeeperNewSize, dedicatedZookeeper); isAllChange {
		nodeSize = newSize
		nodePurpose = nil
	} else {
		newSize, purpose, err := getSingleChangedKafkaSizeAndPurpose(brokerNewSize, zookeeperNewSize, dedicatedZookeeper)
		if err != nil {
			return nil, nil, err
		}
		nodeSize = newSize
		nodePurpose = &purpose
	}

	log.Printf("[INFO] Resizing Kafka cluster. nodePurpose: %s, newSize: %s", nodePurpose, nodeSize)
	res, err := client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, nodeSize, nodePurpose)
	if err != nil {
		return nil, nil, fmt.Errorf("[Error] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return &cluster.DataCentres[0].ID, res, nil
}

func doLegacyCassandraClusterResize(client *APIClient, cluster *Cluster, d *schema.ResourceData) (*string, *ClusterDataCenterResizeResponse, error) {
	before, after := d.GetChange("node_size")
	regex := regexp.MustCompile(`resizeable-(small|large)`)
	oldNodeClass := regex.FindString(before.(string))
	newNodeClass := regex.FindString(after.(string))

	isNotResizable := oldNodeClass == ""
	isNotSameSizeClass := newNodeClass != oldNodeClass
	if isNotResizable || isNotSameSizeClass {
		return nil, nil, fmt.Errorf("[Error] Cannot resize nodes from %s to %s", before, after)
	}

	res, err := client.ResizeCluster(cluster.ID, cluster.DataCentres[0].ID, after.(string), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("[Error] Error resizing cluster %s with error %s", cluster.ID, err)
	}
	return &cluster.DataCentres[0].ID, res, nil
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

	clusterProvider := make(map[string]interface{}, 0)
	mapstructure.Decode(cluster.Provider[0], &clusterProvider)
	processedClusterProvider := processProvider(d, clusterProvider)
	d.Set("cluster_provider", processedClusterProvider)

	bundles, err := getBundlesFromCluster(cluster)
	if err != nil {
		return err
	}

	if err := d.Set("bundle", bundles); err != nil {
		return fmt.Errorf("[Error] Error reading cluster: %s", err)
	}

	nodeSize := ""
	/*
	*  Ideally, we would like this information to be coming directly from the API cluster status.
	*  Hence, this is a slightly hacky way of ignoring zookeeper node sizes (Kafka bundles specific).
	 */
	for _, node := range cluster.DataCentres[0].Nodes {
		nodeSize = node.Size
		if !strings.HasPrefix(nodeSize, "zk-") {
			break
		}
	}

	//loop over each data centre to determine nodes and racks for rack allocation
	nodeCount := 0
	rackList := make([]string, 0)
	for _, dataCentre := range cluster.DataCentres {
		for _, node := range dataCentre.Nodes {
			if !strings.HasPrefix(node.Size, "zk-") {
				nodeCount += 1
			}
			rackList = appendIfMissing(rackList, node.Rack)
		}
	}
	rackCount := len(rackList)
	nodesPerRack := nodeCount / rackCount

	rackAllocation := make(map[string]interface{}, 0)
	rackAllocation["number_of_racks"] = strconv.Itoa(rackCount)
	rackAllocation["nodes_per_rack"] = strconv.Itoa(nodesPerRack)

	if err := d.Set("rack_allocation", rackAllocation); err != nil {
		return fmt.Errorf("[Error] Error reading cluster, rack allocation could not be derived: %s", err)
	}
	if len(cluster.DataCentres[0].ResizeTargetNodeSize) > 0 {
		nodeSize = cluster.DataCentres[0].ResizeTargetNodeSize
	}
	if cluster.BundleType != "ELASTICSEARCH" {
		d.Set("node_size", nodeSize)
	}
	d.Set("data_centre", cluster.DataCentres[0].Name)
	d.Set("sla_tier", strings.ToUpper(cluster.SlaTier))
	d.Set("cluster_network", cluster.DataCentres[0].CdcNetwork)
	d.Set("private_network_cluster", cluster.DataCentres[0].PrivateIPOnly)
	d.Set("pci_compliant_cluster", cluster.PciCompliance == "ENABLED")

	azList := make([]string, 0)
	publicContactPointList := make([]string, 0)
	privateContactPointList := make([]string, 0)

	for _, dataCentre := range cluster.DataCentres {
		for _, node := range dataCentre.Nodes {
			if !stringInSlice(node.Rack, azList) {
				if !strings.HasPrefix(node.Size, "zk-") {
					azList = appendIfMissing(azList, node.Rack)
					privateContactPointList = appendIfMissing(privateContactPointList, node.PrivateAddress)
					publicContactPointList = appendIfMissing(publicContactPointList, node.PublicAddress)
				}
			}
		}
	}

	if !cluster.DataCentres[0].PrivateIPOnly && len(publicContactPointList) > 0 {
		err = d.Set("public_contact_point", publicContactPointList)
	} else {
		err = d.Set("public_contact_point", nil)
	}

	if len(privateContactPointList) > 0 {
		err = d.Set("private_contact_point", privateContactPointList)
	} else {
		err = d.Set("public_contact_point", nil)
	}

	toCheck := [2]string{"cluster_provider", "rack_allocation"}
	for _, changing := range toCheck {

		if !d.HasChange(changing) {
			continue
		}
		_, after := d.GetChange(changing)
		d.Set(changing, after)
	}

	log.Printf("[INFO] Fetched cluster %s info from the remote server.", cluster.ID)
	return nil
}

func getBundlesFromCluster(cluster *Cluster) ([]map[string]interface{}, error) {
	baseBundle := make(map[string]interface{}, 3)
	baseBundle["bundle"] = cluster.BundleType

	baseBundleOptions := make(map[string]interface{}, 0)
	err := mapstructure.Decode(cluster.BundleOption, &baseBundleOptions)
	if err != nil {
		return nil, fmt.Errorf("[Error] Error decoding bundles option: %s", err)
	}

	convertedBundleOptions := dereferencePointerInStruct(baseBundleOptions)

	baseBundle["options"] = convertedBundleOptions
	baseBundle["version"] = cluster.BundleVersion

	bundles := make([]map[string]interface{}, 0)
	bundles = append(bundles, baseBundle)

	addonBundles := cluster.AddonBundles

	if addonBundles == nil {
		return nil, nil
	}

	sort.Slice(addonBundles, func(i, j int) bool { return addonBundles[i]["bundle"].(string) > addonBundles[j]["bundle"].(string) })
	for _, addonBundle := range addonBundles {
		if len(addonBundle) != 0 {
			bundles = append(bundles, addonBundle)
		}
	}

	return bundles, nil
}

func dereferencePointerInStruct(data map[string]interface{}) map[string]interface{} {
	for k, v := range data {
		// Terraform expects strings for everything
		// This block iterates through the bundle options map and checks for an interface of a pointer
		// For each interface{*type} value, it changes: interface{*type} -> *type -> type -> interface{type} -> String
		// For a non-pointer, it directly formats to a string
		valueOfV := reflect.ValueOf(v)
		if valueOfV.Kind() == reflect.Ptr {
			data[k] = fmt.Sprintf("%v", reflect.Indirect(valueOfV.Elem()).Interface())
		} else {
			data[k] = fmt.Sprintf("%v", valueOfV)
		}
	}

	return data
}

func processProvider(d *schema.ResourceData, data map[string]interface{}) (newData map[string]interface{}) {
	// This is used to ignore values that are not set in the resource
	// Otherwise terraform will store the default and generated values (e.g. Provider Account Name)
	// into the state and result a diff between the state and resource (plan)
	newData = make(map[string]interface{})
	for k, v := range data {
		resource := d.Get(fmt.Sprintf("cluster_provider.%s", k))
		// Store the required field "name"
		// There doesn't seem to be a way to programmatically get the "required" schemas with ResourceData
		if resource != "" || k == "name" {
			newData[k] = v
		}
	}

	return newData
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

func resourceClusterStateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	clusterId := d.Id()
	d.Set("cluster_id", clusterId)
	return []*schema.ResourceData{d}, nil
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
		"APACHE_ZOOKEEPER",
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
