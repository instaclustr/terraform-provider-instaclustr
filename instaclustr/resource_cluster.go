package instaclustr

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
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
				Required: true,
			},

			"data_centre": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"data_centres": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
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

	var createData = CreateRequest{
		ClusterName:           d.Get("cluster_name").(string),
		Bundles:               bundles,
		Provider:              clusterProvider,
		SlaTier:               d.Get("sla_tier").(string),
		NodeSize:              d.Get("node_size").(string),
		PrivateNetworkCluster: fmt.Sprintf("%v", d.Get("private_network_cluster")),
		PCICompliantCluster:   fmt.Sprintf("%v", d.Get("pci_compliant_cluster")),
	}

	dataCentre := d.Get("data_centre").(string)

	if dataCentre != "" {
		clusterNetwork := d.Get("cluster_network").(string)
		createData.DataCentre = dataCentre
		createData.ClusterNetwork = clusterNetwork
	} else {
		dataCentres, err := getDataCentres(d)
		if err != nil {
			return formatCreateErrMsg(err)
		}
		createData.DataCentres = dataCentres
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

	return waitForClusterStateAndDoUpdate(client, waitForClusterState, bundleConfig, kafkaRestProxyUserPassword, kafkaSchemaRegistryUserPassword, d, id)
}

func waitForClusterStateAndDoUpdate(client *APIClient,
	waitForClusterState string,
	bundleConfig BundleConfig,
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

func resourceClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(true)
	// currently only cluster resize, kafka-schema-registry user password update and kafka-rest-proxy user password update are supported

	client := meta.(*Config).Client
	clusterID := d.Get("cluster_id").(string)

	clusterResize := d.HasChange("node_size")
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

	if clusterResize {
		//resizing the cluster (i.e, upgrading from one node size to another node size)
		err = doClusterResize(client, clusterID, d)
		if err != nil {
			return fmt.Errorf("[Error] Error resizing the cluster : %s", err)
		}
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

	nodeCount := 0
	rackList := make([]string, 0)
	for _, node := range cluster.DataCentres[0].Nodes {
		if !strings.HasPrefix(node.Size, "zk-") {
			nodeCount += 1
		}
		rackList = appendIfMissing(rackList, node.Rack)
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

	d.Set("data_centre", cluster.DataCentre)
	dataCentres, err := getDataCentresFromCluster(cluster, d)
	if err != nil {
		return err
	}
	if err := d.Set("data_centres", dataCentres); err != nil {
		return fmt.Errorf("[Error] Error reading cluster, data centres could not be derived: %s", err)
	}
	d.Set("node_size", nodeSize)
	d.Set("sla_tier", strings.ToUpper(cluster.SlaTier))
	if cluster.DataCentre != "" {
		d.Set("cluster_network", cluster.DataCentres[0].CdcNetwork)
	}
	d.Set("private_network_cluster", cluster.DataCentres[0].PrivateIPOnly)
	d.Set("pci_compliant_cluster", cluster.PciCompliance == "ENABLED")

	if len(cluster.DataCentres[0].Nodes[0].PublicAddress) != 0 {
		err = d.Set("public_contact_point", cluster.DataCentres[0].Nodes[0].PublicAddress)
	}

	if len(cluster.DataCentres[0].Nodes[0].PrivateAddress) != 0 {
		err = d.Set("private_contact_point", cluster.DataCentres[0].Nodes[0].PrivateAddress)
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
	if cluster.AddonBundles != nil {
		for _, addonBundle := range cluster.AddonBundles {
			if addonBundle != nil {
				bundles = append(bundles, addonBundle)
			}
		}
	}

	return bundles, nil
}

func getDataCentresFromCluster(cluster *Cluster, d *schema.ResourceData) ([]map[string]string, error) {
	dataCentres := make([]map[string]string, 0)
	stateDataCentres, err := getDataCentres(d)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to get data centres from schema.ResourceData: %s", err)
	}

	// make sure the order of data centres is correct.
	for _, stateDataCentre := range stateDataCentres {
		if cluster.DataCentres != nil {
			for _, dataCentre := range cluster.DataCentres {
				if dataCentre.CdcNetwork == stateDataCentre.Network {
					dataCentreMap := map[string]string{
						"data_centre_region": dataCentre.Name,
						"network":            dataCentre.CdcNetwork,
					}
					dataCentres = append(dataCentres, dataCentreMap)
				}
			}
		}
	}
	return dataCentres, nil
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

func getDataCentres(d *schema.ResourceData) ([]DataCentre, error) {
	dataCentres := make([]DataCentre, 0)
	for _, inDataCentre := range d.Get("data_centres").([]interface{}) {
		var dataCentre DataCentre
		err := mapstructure.WeakDecode(inDataCentre.(map[string]interface{}), &dataCentre)
		if err != nil {
			return nil, err
		}
		dataCentres = append(dataCentres, dataCentre)
	}
	return dataCentres, nil
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
